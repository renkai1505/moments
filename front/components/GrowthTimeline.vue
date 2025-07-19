<template>
  <div class="relative">
    <!-- 加载状态 -->
    <div v-if="loading" class="text-center py-8">
      <UIcon name="i-heroicons-arrow-path" class="w-8 h-8 animate-spin text-gray-400" />
      <p class="text-gray-500 mt-2">加载中...</p>
    </div>

    <!-- 时间轴内容 -->
    <div v-else-if="timeline.length > 0" class="space-y-8">
      <div v-for="(item, index) in timeline" :key="item.date" class="relative">
        <!-- 时间轴线 -->
        <div v-if="index < timeline.length - 1" 
             class="absolute left-6 top-16 w-0.5 h-full bg-gray-200 dark:bg-neutral-600"></div>
        
        <!-- 日期标签 -->
        <div class="flex items-center mb-6">
          <div class="w-12 h-12 bg-primary-500 rounded-full flex items-center justify-center text-white font-bold z-10">
            {{ formatDayMonth(item.date) }}
          </div>
          <div class="ml-4">
            <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
              {{ formatFullDate(item.date) }}
            </h3>
            <p class="text-sm text-gray-500">
              {{ item.records.length }}条记录
            </p>
          </div>
        </div>

        <!-- 该日期的记录 -->
        <div class="ml-16 space-y-4">
          <div v-for="record in item.records" :key="record.id" 
               class="bg-white dark:bg-neutral-800 rounded-lg shadow-sm border border-gray-200 dark:border-neutral-700 p-4">
            
            <!-- 记录标题和类型 -->
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center space-x-3">
                <div class="w-8 h-8 rounded-full flex items-center justify-center text-sm"
                     :class="getTypeColor(record.recordType)">
                  {{ getTypeIcon(record.recordType) }}
                </div>
                <div>
                  <h4 class="font-semibold text-gray-900 dark:text-white">
                    {{ record.title || '成长记录' }}
                  </h4>
                  <p class="text-xs text-gray-500">
                    {{ formatTime(record.recordDate || record.createdAt) }}
                  </p>
                </div>
              </div>
              <UBadge :color="getTypeBadgeColor(record.recordType)" variant="soft" size="xs">
                {{ getTypeLabel(record.recordType) }}
              </UBadge>
            </div>

            <!-- 记录内容 -->
            <p class="text-gray-700 dark:text-gray-300 text-sm mb-3 leading-relaxed">
              {{ record.content }}
            </p>

            <!-- 数据信息 -->
            <div v-if="record.height || record.weight || record.location || record.mood" 
                 class="grid grid-cols-2 md:grid-cols-4 gap-3 mb-3">
              <div v-if="record.height" class="flex items-center space-x-2 text-sm">
                <UIcon name="i-heroicons-arrows-up-down" class="w-4 h-4 text-blue-500" />
                <span class="text-blue-600 font-medium">{{ record.height }}cm</span>
              </div>
              <div v-if="record.weight" class="flex items-center space-x-2 text-sm">
                <UIcon name="i-heroicons-scale" class="w-4 h-4 text-green-500" />
                <span class="text-green-600 font-medium">{{ record.weight }}kg</span>
              </div>
              <div v-if="record.location" class="flex items-center space-x-2 text-sm">
                <UIcon name="i-heroicons-map-pin" class="w-4 h-4 text-gray-500" />
                <span class="text-gray-600">{{ record.location }}</span>
              </div>
              <div v-if="record.mood" class="flex items-center space-x-2 text-sm">
                <UIcon name="i-heroicons-face-smile" class="w-4 h-4 text-yellow-500" />
                <span class="text-yellow-600">{{ record.mood }}</span>
              </div>
            </div>

            <!-- 里程碑 -->
            <div v-if="record.milestone" 
                 class="bg-gradient-to-r from-yellow-50 to-orange-50 dark:from-yellow-900/20 dark:to-orange-900/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-3 mb-3">
              <div class="flex items-center space-x-2 mb-2">
                <UIcon name="i-heroicons-trophy" class="w-5 h-5 text-yellow-600" />
                <span class="font-medium text-yellow-800 dark:text-yellow-200">重要里程碑</span>
              </div>
              <p class="text-yellow-700 dark:text-yellow-300 text-sm">{{ record.milestone }}</p>
            </div>

            <!-- 图片网格 -->
            <div v-if="record.imgConfigs && record.imgConfigs.length > 0" class="mb-3">
              <div class="grid grid-cols-3 md:grid-cols-4 gap-2">
                <img 
                  v-for="(img, imgIndex) in record.imgConfigs.slice(0, 6)" 
                  :key="imgIndex"
                  :src="img.thumbUrl || img.url" 
                  :alt="`图片${imgIndex + 1}`"
                  class="aspect-square object-cover rounded-lg cursor-pointer hover:opacity-80 transition-opacity"
                  @click="showImageModal(record.imgConfigs, imgIndex)"
                />
                <div v-if="record.imgConfigs.length > 6" 
                     class="aspect-square bg-gray-100 dark:bg-neutral-700 rounded-lg flex items-center justify-center text-gray-500 text-sm cursor-pointer hover:bg-gray-200 dark:hover:bg-neutral-600 transition-colors"
                     @click="showImageModal(record.imgConfigs, 6)">
                  +{{ record.imgConfigs.length - 6 }}
                </div>
              </div>
            </div>

            <!-- 标签 -->
            <div v-if="record.tags" class="flex flex-wrap gap-1 mb-3">
              <UBadge 
                v-for="tag in getTags(record.tags)" 
                :key="tag" 
                color="gray" 
                variant="soft" 
                size="xs"
              >
                {{ tag }}
              </UBadge>
            </div>

            <!-- 操作按钮 -->
            <div class="flex items-center justify-end space-x-2 pt-3 border-t border-gray-100 dark:border-neutral-600">
              <UButton 
                icon="i-heroicons-eye" 
                variant="ghost" 
                size="sm"
                @click="navigateTo(`/growth/${record.id}`)"
              >
                查看
              </UButton>
              <UButton 
                icon="i-heroicons-pencil" 
                variant="ghost" 
                size="sm"
                @click="navigateTo(`/growth/edit/${record.id}`)"
              >
                编辑
              </UButton>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-else class="text-center py-16">
      <UIcon name="i-heroicons-clock" class="w-16 h-16 text-gray-300 mx-auto mb-4" />
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">暂无成长记录</h3>
      <p class="text-gray-500 mb-6">开始记录宝贝的成长时光吧！</p>
      <UButton 
        color="primary" 
        @click="navigateTo(`/growth/new?childId=${childId}`)"
      >
        <UIcon name="i-heroicons-plus" class="mr-2" />
        添加第一条记录
      </UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { TimelineItemVO } from "~/types";
import { format } from 'date-fns';
import { zhCN } from 'date-fns/locale';

interface Props {
  childId: number;
}

const props = defineProps<Props>();

const loading = ref(true);
const timeline = ref<TimelineItemVO[]>([]);

onMounted(async () => {
  await loadTimeline();
});

const loadTimeline = async () => {
  try {
    loading.value = true;
    const res = await useMyFetch<TimelineItemVO[]>(`/growth/timeline/${props.childId}`);
    timeline.value = res || [];
  } catch (error) {
    console.error('加载时间轴失败:', error);
  } finally {
    loading.value = false;
  }
};

const getTypeIcon = (type: string) => {
  const icons = {
    'growth': '📏',
    'health': '🩺',
    'study': '📚',
    'play': '🎮',
    'milestone': '🏆'
  };
  return icons[type] || '📝';
};

const getTypeColor = (type: string) => {
  const colors = {
    'growth': 'bg-blue-100 text-blue-600',
    'health': 'bg-red-100 text-red-600',
    'study': 'bg-green-100 text-green-600',
    'play': 'bg-yellow-100 text-yellow-600',
    'milestone': 'bg-purple-100 text-purple-600'
  };
  return colors[type] || 'bg-gray-100 text-gray-600';
};

const getTypeBadgeColor = (type: string) => {
  const colors = {
    'growth': 'blue',
    'health': 'red',
    'study': 'green',
    'play': 'yellow',
    'milestone': 'purple'
  };
  return colors[type] || 'gray';
};

const getTypeLabel = (type: string) => {
  const labels = {
    'growth': '成长发育',
    'health': '健康记录',
    'study': '学习记录',
    'play': '游戏玩耍',
    'milestone': '重要时刻'
  };
  return labels[type] || '其他';
};

const getTags = (tagsStr: string) => {
  return tagsStr ? tagsStr.split(',').filter(tag => tag.trim()) : [];
};

const formatDayMonth = (dateStr: string) => {
  if (!dateStr) return '';
  try {
    return format(new Date(dateStr), 'dd');
  } catch {
    return '';
  }
};

const formatFullDate = (dateStr: string) => {
  if (!dateStr) return '';
  try {
    return format(new Date(dateStr), 'yyyy年MM月dd日 EEEE', { locale: zhCN });
  } catch {
    return dateStr;
  }
};

const formatTime = (dateStr: string) => {
  if (!dateStr) return '';
  try {
    return format(new Date(dateStr), 'HH:mm');
  } catch {
    return '';
  }
};

const showImageModal = (images: any[], startIndex: number) => {
  // TODO: 实现图片查看器
  console.log('显示图片', images, startIndex);
};
</script>

<style scoped>
.aspect-square {
  aspect-ratio: 1 / 1;
}
</style>