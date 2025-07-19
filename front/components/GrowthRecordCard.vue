<template>
  <div class="bg-white dark:bg-neutral-800 rounded-lg shadow-sm border border-gray-200 dark:border-neutral-700 p-4 hover:shadow-md transition-shadow">
    <div class="flex items-start space-x-4">
      <!-- 记录类型图标 -->
      <div class="flex-shrink-0">
        <div class="w-12 h-12 rounded-full flex items-center justify-center text-white text-xl"
             :class="getTypeColor(record.recordType)">
          {{ getTypeIcon(record.recordType) }}
        </div>
      </div>

      <!-- 记录内容 -->
      <div class="flex-1 min-w-0">
        <div class="flex items-center justify-between mb-2">
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white line-clamp-1">
            {{ record.title || '成长记录' }}
          </h3>
          <div class="flex items-center space-x-2 text-sm text-gray-500">
            <span>{{ formatDate(record.recordDate || record.createdAt) }}</span>
            <UBadge :color="getTypeBadgeColor(record.recordType)" variant="soft" size="xs">
              {{ getTypeLabel(record.recordType) }}
            </UBadge>
          </div>
        </div>

        <!-- 内容预览 -->
        <p class="text-gray-700 dark:text-gray-300 text-sm mb-3 line-clamp-2">
          {{ record.content }}
        </p>

        <!-- 特殊信息 -->
        <div class="flex items-center flex-wrap gap-3 mb-3 text-sm">
          <div v-if="record.height" class="flex items-center space-x-1 text-blue-600">
            <UIcon name="i-heroicons-arrows-up-down" class="w-4 h-4" />
            <span>{{ record.height }}cm</span>
          </div>
          <div v-if="record.weight" class="flex items-center space-x-1 text-green-600">
            <UIcon name="i-heroicons-scale" class="w-4 h-4" />
            <span>{{ record.weight }}kg</span>
          </div>
          <div v-if="record.location" class="flex items-center space-x-1 text-gray-600">
            <UIcon name="i-heroicons-map-pin" class="w-4 h-4" />
            <span>{{ record.location }}</span>
          </div>
          <div v-if="record.mood" class="flex items-center space-x-1 text-yellow-600">
            <UIcon name="i-heroicons-face-smile" class="w-4 h-4" />
            <span>{{ record.mood }}</span>
          </div>
        </div>

        <!-- 里程碑信息 -->
        <div v-if="record.milestone" class="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-3 mb-3">
          <div class="flex items-center space-x-2">
            <UIcon name="i-heroicons-trophy" class="w-5 h-5 text-yellow-600" />
            <span class="font-medium text-yellow-800 dark:text-yellow-200">重要里程碑</span>
          </div>
          <p class="text-yellow-700 dark:text-yellow-300 text-sm mt-1">{{ record.milestone }}</p>
        </div>

        <!-- 图片预览 -->
        <div v-if="record.imgConfigs && record.imgConfigs.length > 0" class="mb-3">
          <div class="flex space-x-2 overflow-x-auto">
            <img 
              v-for="(img, index) in record.imgConfigs.slice(0, 3)" 
              :key="index"
              :src="img.thumbUrl || img.url" 
              :alt="`图片${index + 1}`"
              class="w-16 h-16 object-cover rounded-lg flex-shrink-0 cursor-pointer"
              @click="showImageModal(record.imgConfigs, index)"
            />
            <div v-if="record.imgConfigs.length > 3" 
                 class="w-16 h-16 bg-gray-100 dark:bg-neutral-700 rounded-lg flex items-center justify-center text-gray-500 text-sm cursor-pointer"
                 @click="showImageModal(record.imgConfigs, 3)">
              +{{ record.imgConfigs.length - 3 }}
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

        <!-- 底部操作栏 -->
        <div class="flex items-center justify-between pt-3 border-t border-gray-100 dark:border-neutral-700">
          <div class="flex items-center space-x-4 text-sm text-gray-500">
            <div class="flex items-center space-x-1">
              <UAvatar :src="record.parent?.avatarUrl" :alt="record.parent?.nickname" size="xs" />
              <span>{{ record.parent?.nickname }}</span>
            </div>
            <span>{{ record.child?.name }}</span>
          </div>
          
          <div class="flex items-center space-x-2">
            <UButton 
              icon="i-heroicons-eye" 
              variant="ghost" 
              size="sm"
              @click="navigateTo(`/growth/${record.id}`)"
            />
            <UButton 
              icon="i-heroicons-pencil" 
              variant="ghost" 
              size="sm"
              @click="navigateTo(`/growth/edit/${record.id}`)"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { GrowthRecordVO } from "~/types";
import { format } from 'date-fns';

interface Props {
  record: GrowthRecordVO;
}

const props = defineProps<Props>();

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
    'growth': 'bg-blue-500',
    'health': 'bg-red-500',
    'study': 'bg-green-500',
    'play': 'bg-yellow-500',
    'milestone': 'bg-purple-500'
  };
  return colors[type] || 'bg-gray-500';
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

const formatDate = (dateStr: string) => {
  if (!dateStr) return '';
  try {
    return format(new Date(dateStr), 'MM月dd日');
  } catch {
    return dateStr;
  }
};

const showImageModal = (images: any[], startIndex: number) => {
  // TODO: 实现图片查看器
  console.log('显示图片', images, startIndex);
};
</script>

<style scoped>
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>