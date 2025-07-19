<template>
  <div v-if="child" class="min-h-screen bg-gray-50 dark:bg-neutral-900">
    <!-- 头部导航 -->
    <div class="bg-white dark:bg-neutral-800 shadow-sm">
      <div class="max-w-4xl mx-auto px-4 py-3">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <UButton 
              icon="i-heroicons-arrow-left" 
              variant="ghost" 
              @click="navigateTo('/')"
            />
            <h1 class="text-xl font-bold">{{ child.name }}的成长档案</h1>
          </div>
          <div class="flex items-center space-x-2">
            <UButton 
              icon="i-heroicons-plus" 
              color="primary" 
              @click="navigateTo(`/growth/new?childId=${child.id}`)"
            >
              添加记录
            </UButton>
            <UDropdown 
              :items="[
                [{ label: '编辑档案', icon: 'i-heroicons-pencil', click: () => navigateTo(`/child/edit/${child.id}`) }],
                [{ label: '统计报告', icon: 'i-heroicons-chart-bar', click: () => showStats = true }]
              ]"
            >
              <UButton icon="i-heroicons-ellipsis-vertical" variant="ghost" />
            </UDropdown>
          </div>
        </div>
      </div>
    </div>

    <!-- 儿童信息卡片 -->
    <div class="max-w-4xl mx-auto px-4 py-6">
      <div class="bg-white dark:bg-neutral-800 rounded-lg shadow-lg overflow-hidden mb-6">
        <!-- 封面图 -->
        <div class="relative h-48 bg-gradient-to-br from-blue-400 to-purple-500">
          <img v-if="child.coverUrl" :src="child.coverUrl" alt="封面" class="w-full h-full object-cover">
          <div class="absolute inset-0 bg-black bg-opacity-20"></div>
          <div class="absolute bottom-6 left-6 text-white">
            <h2 class="text-3xl font-bold mb-2">{{ child.name }}</h2>
            <p class="text-lg opacity-90">{{ child.nickname || child.name }}</p>
            <div class="flex items-center space-x-4 mt-2 text-sm">
              <span>{{ child.age }}岁{{ child.ageInDays % 365 }}天</span>
              <span>•</span>
              <span>{{ child.gender === 'M' ? '👦 男孩' : '👧 女孩' }}</span>
              <span>•</span>
              <span>{{ formatDate(child.birthDate) }}出生</span>
            </div>
          </div>
          <div class="absolute top-6 right-6">
            <UAvatar :src="child.avatarUrl" :alt="child.name" size="lg" />
          </div>
        </div>

        <!-- 基本信息 -->
        <div class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
            <div class="text-center p-4 bg-gray-50 dark:bg-neutral-700 rounded-lg">
              <div class="text-2xl font-bold text-blue-600">{{ child.height || '--' }}</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">身高(cm)</div>
            </div>
            <div class="text-center p-4 bg-gray-50 dark:bg-neutral-700 rounded-lg">
              <div class="text-2xl font-bold text-green-600">{{ child.weight || '--' }}</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">体重(kg)</div>
            </div>
            <div class="text-center p-4 bg-gray-50 dark:bg-neutral-700 rounded-lg">
              <div class="text-2xl font-bold text-red-600">{{ child.bloodType || '--' }}</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">血型</div>
            </div>
            <div class="text-center p-4 bg-gray-50 dark:bg-neutral-700 rounded-lg">
              <div class="text-2xl font-bold text-purple-600">{{ child.ageInDays }}</div>
              <div class="text-sm text-gray-600 dark:text-gray-400">出生天数</div>
            </div>
          </div>

          <div v-if="child.description" class="mb-4">
            <h3 class="text-lg font-semibold mb-2">简介</h3>
            <p class="text-gray-700 dark:text-gray-300">{{ child.description }}</p>
          </div>

          <div v-if="child.hobbies" class="mb-4">
            <h3 class="text-lg font-semibold mb-2">兴趣爱好</h3>
            <p class="text-gray-700 dark:text-gray-300">{{ child.hobbies }}</p>
          </div>
        </div>
      </div>

      <!-- 功能导航 -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
        <UButton 
          class="h-20 flex flex-col items-center justify-center"
          variant="outline"
          @click="activeTab = 'timeline'"
        >
          <UIcon name="i-heroicons-clock" class="text-2xl mb-1" />
          <span class="text-sm">成长时间轴</span>
        </UButton>
        <UButton 
          class="h-20 flex flex-col items-center justify-center"
          variant="outline"
          @click="activeTab = 'records'"
        >
          <UIcon name="i-heroicons-document-text" class="text-2xl mb-1" />
          <span class="text-sm">成长记录</span>
        </UButton>
        <UButton 
          class="h-20 flex flex-col items-center justify-center"
          variant="outline"
          @click="activeTab = 'milestones'"
        >
          <UIcon name="i-heroicons-trophy" class="text-2xl mb-1" />
          <span class="text-sm">里程碑</span>
        </UButton>
        <UButton 
          class="h-20 flex flex-col items-center justify-center"
          variant="outline"
          @click="showStats = true"
        >
          <UIcon name="i-heroicons-chart-bar" class="text-2xl mb-1" />
          <span class="text-sm">成长统计</span>
        </UButton>
      </div>

      <!-- 内容区域 -->
      <div class="bg-white dark:bg-neutral-800 rounded-lg shadow-lg">
        <!-- 成长时间轴 -->
        <div v-if="activeTab === 'timeline'" class="p-6">
          <h3 class="text-xl font-bold mb-6">📖 成长时间轴</h3>
          <GrowthTimeline :childId="child.id" />
        </div>

        <!-- 成长记录列表 -->
        <div v-else-if="activeTab === 'records'" class="p-6">
          <div class="flex items-center justify-between mb-6">
            <h3 class="text-xl font-bold">📝 成长记录</h3>
            <div class="flex items-center space-x-2">
              <USelectMenu 
                v-model="recordTypeFilter"
                :options="recordTypeOptions"
                placeholder="筛选类型"
                size="sm"
              />
              <UButton 
                icon="i-heroicons-plus" 
                size="sm"
                @click="navigateTo(`/growth/new?childId=${child.id}`)"
              >
                添加记录
              </UButton>
            </div>
          </div>
          <GrowthRecordList :childId="child.id" :recordType="recordTypeFilter" />
        </div>

        <!-- 里程碑 -->
        <div v-else-if="activeTab === 'milestones'" class="p-6">
          <h3 class="text-xl font-bold mb-6">🏆 重要里程碑</h3>
          <MilestoneList :childId="child.id" />
        </div>
      </div>
    </div>

    <!-- 统计弹窗 -->
    <GrowthStatsModal 
      v-if="showStats" 
      :childId="child.id" 
      @close="showStats = false" 
    />
  </div>
  <div v-else class="flex items-center justify-center h-screen">
    <div class="text-center">
      <UIcon name="i-heroicons-face-frown" class="text-6xl text-gray-400 mb-4" />
      <h2 class="text-xl font-bold text-gray-600 dark:text-gray-400 mb-2">儿童档案不存在</h2>
      <p class="text-gray-500 mb-4">该儿童档案可能已被删除或您没有访问权限</p>
      <UButton @click="navigateTo('/')">返回首页</UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ChildVO } from "~/types";
import { format } from 'date-fns';

const route = useRoute()
const childId = Number(route.params.id)

const child = ref<ChildVO | null>(null)
const activeTab = ref('timeline')
const showStats = ref(false)
const recordTypeFilter = ref('')

const recordTypeOptions = [
  { label: '全部', value: '' },
  { label: '成长发育', value: 'growth' },
  { label: '健康记录', value: 'health' },
  { label: '学习记录', value: 'study' },
  { label: '游戏玩耍', value: 'play' },
  { label: '重要时刻', value: 'milestone' }
]

onMounted(async () => {
  await loadChild()
})

const loadChild = async () => {
  try {
    const res = await useMyFetch<ChildVO>(`/child/${childId}`)
    child.value = res
  } catch (error) {
    console.error('加载儿童档案失败:', error)
  }
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return ''
  try {
    return format(new Date(dateStr), 'yyyy年MM月dd日')
  } catch {
    return dateStr
  }
}

// SEO
useSeoMeta({
  title: () => child.value ? `${child.value.name}的成长档案` : '儿童档案',
  description: () => child.value ? `${child.value.name}的成长记录和重要时刻` : '查看儿童成长档案'
})
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>