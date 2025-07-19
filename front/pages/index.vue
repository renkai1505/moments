<template>
  <Header v-bind:user="currentUser"/>
  <div class="max-w-4xl mx-auto px-4 py-6">
    <!-- 儿童列表 -->
    <div v-if="children.length > 0" class="mb-8">
      <h2 class="text-2xl font-bold mb-6 text-center">🌟 儿童相册</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <div v-for="child in children" :key="child.id" 
             class="bg-white dark:bg-neutral-800 rounded-lg shadow-lg overflow-hidden cursor-pointer hover:shadow-xl transition-shadow"
             @click="navigateToChild(child.id)">
          <div class="relative h-48 bg-gradient-to-br from-blue-400 to-purple-500">
            <img v-if="child.coverUrl" :src="child.coverUrl" alt="封面" class="w-full h-full object-cover">
            <div class="absolute bottom-4 left-4 text-white">
              <h3 class="text-xl font-bold">{{ child.name }}</h3>
              <p class="text-sm opacity-90">{{ child.nickname || child.name }}</p>
              <p class="text-xs opacity-75">{{ child.age }}岁 • {{ child.ageInDays }}天</p>
            </div>
            <div class="absolute top-4 right-4">
              <UAvatar :src="child.avatarUrl" :alt="child.name" size="sm" />
            </div>
          </div>
          <div class="p-4">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm text-gray-600 dark:text-gray-400">
                {{ child.gender === 'M' ? '👦 男孩' : '👧 女孩' }}
              </span>
              <UBadge color="blue" variant="soft" size="xs">
                {{ formatDate(child.birthDate) }}
              </UBadge>
            </div>
            <p class="text-sm text-gray-700 dark:text-gray-300 line-clamp-2">
              {{ child.description || '这是一个可爱的宝贝~' }}
            </p>
          </div>
        </div>
      </div>
    </div>

    <!-- 如果没有儿童，显示引导创建 -->
    <div v-else class="text-center py-16">
      <div class="mb-8">
        <div class="text-6xl mb-4">👶</div>
        <h2 class="text-2xl font-bold mb-4">欢迎来到儿童相册</h2>
        <p class="text-gray-600 dark:text-gray-400 mb-8">
          开始记录宝贝的成长时光，创建第一个儿童档案吧！
        </p>
        <UButton color="primary" size="lg" @click="navigateTo('/child/new')">
          <UIcon name="i-heroicons-plus" class="mr-2" />
          创建儿童档案
        </UButton>
      </div>
    </div>

    <!-- 最新成长记录 -->
    <div v-if="latestRecords.length > 0" class="mb-8">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-xl font-bold">📖 最新成长记录</h2>
        <UButton variant="ghost" size="sm" @click="navigateTo('/growth')">
          查看全部
          <UIcon name="i-heroicons-arrow-right" class="ml-1" />
        </UButton>
      </div>
      <div class="space-y-4">
        <GrowthRecordCard v-for="record in latestRecords" :key="record.id" :record="record" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ChildVO, GrowthRecordVO, SysConfigVO, UserVO } from "~/types";
import { format } from 'date-fns';

const currentUser = useState<UserVO>('userinfo')
const sysConfig = useState<SysConfigVO>('sysConfig')

const children = ref<ChildVO[]>([])
const latestRecords = ref<GrowthRecordVO[]>([])

onMounted(async () => {
  await loadChildren()
  await loadLatestRecords()
})

const loadChildren = async () => {
  try {
    const res = await useMyFetch<ChildVO[]>('/child/list')
    children.value = res || []
  } catch (error) {
    console.error('加载儿童列表失败:', error)
  }
}

const loadLatestRecords = async () => {
  try {
    const res = await useMyFetch<{
      list: GrowthRecordVO[]
      total: number
      hasNext: boolean
    }>('/growth/list', {
      page: 1,
      size: 5
    })
    latestRecords.value = res.list || []
  } catch (error) {
    console.error('加载最新记录失败:', error)
  }
}

const navigateToChild = (childId: number) => {
  navigateTo(`/child/${childId}`)
}

const formatDate = (dateStr: string) => {
  if (!dateStr) return ''
  try {
    return format(new Date(dateStr), 'yyyy-MM-dd')
  } catch {
    return dateStr
  }
}
</script>

<style scoped>

</style>