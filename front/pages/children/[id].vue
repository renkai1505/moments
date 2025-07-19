<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <Header v-bind:user="currentUser"/>
    
    <div v-if="child" class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- 儿童基本信息 -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 mb-8">
        <div class="flex items-center mb-6">
          <UAvatar 
            :src="child.avatar || '/default-child-avatar.svg'" 
            :alt="child.name"
            size="xl"
            class="mr-6"
          />
          <div class="flex-1">
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
              {{ child.name }}
            </h1>
            <div class="flex items-center space-x-4 text-gray-600 dark:text-gray-400">
              <span>{{ child.gender === '男' ? '男孩' : '女孩' }}</span>
              <span v-if="child.birthDate">
                {{ calculateAge(child.birthDate) }}岁
                ({{ formatDate(child.birthDate) }})
              </span>
              <span v-if="child.bloodType">血型：{{ child.bloodType }}</span>
            </div>
            <p v-if="child.description" class="mt-2 text-gray-600 dark:text-gray-400">
              {{ child.description }}
            </p>
          </div>
          <div class="flex space-x-2">
            <UButton 
              @click="showGrowthModal = true"
              color="green"
              icon="i-heroicons-chart-bar"
            >
              添加成长记录
            </UButton>
            <UButton 
              @click="showMilestoneModal = true"
              color="purple"
              icon="i-heroicons-star"
            >
              添加里程碑
            </UButton>
          </div>
        </div>
      </div>

      <!-- 时间轴导航 -->
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6 mb-8">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-semibold text-gray-900 dark:text-white">成长时间轴</h2>
          <div class="flex space-x-2">
            <USelect
              v-model="selectedYear"
              :options="yearOptions"
              placeholder="选择年份"
              class="w-32"
            />
            <USelect
              v-model="selectedMonth"
              :options="monthOptions"
              placeholder="选择月份"
              class="w-32"
            />
            <UButton 
              @click="clearFilter"
              color="gray"
              variant="soft"
              size="sm"
            >
              清除筛选
            </UButton>
          </div>
        </div>

        <!-- 时间轴 -->
        <div class="relative">
          <div class="absolute left-8 top-0 bottom-0 w-0.5 bg-gray-200 dark:bg-gray-700"></div>
          
          <div class="space-y-8">
            <div 
              v-for="(group, date) in groupedRecords" 
              :key="date"
              class="relative"
            >
              <!-- 日期标题 -->
              <div class="flex items-center mb-4">
                <div class="w-16 h-16 bg-blue-500 rounded-full flex items-center justify-center text-white font-semibold mr-4">
                  {{ formatDay(date) }}
                </div>
                <div>
                  <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    {{ formatDateTitle(date) }}
                  </h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">
                    {{ formatDateFull(date) }}
                  </p>
                </div>
              </div>

              <!-- 记录列表 -->
              <div class="ml-20 space-y-4">
                <div 
                  v-for="record in group" 
                  :key="record.id"
                  class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4 relative"
                >
                  <!-- 成长记录 -->
                  <div v-if="record.type === 'growth'" class="flex items-start">
                    <div class="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center text-white text-sm mr-4 mt-1">
                      <UIcon name="i-heroicons-chart-bar" class="w-4 h-4" />
                    </div>
                    <div class="flex-1">
                      <h4 class="font-medium text-gray-900 dark:text-white mb-2">成长记录</h4>
                      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm">
                        <div v-if="record.height" class="flex items-center">
                          <span class="text-gray-500 dark:text-gray-400 mr-2">身高：</span>
                          <span class="font-medium">{{ record.height }}cm</span>
                        </div>
                        <div v-if="record.weight" class="flex items-center">
                          <span class="text-gray-500 dark:text-gray-400 mr-2">体重：</span>
                          <span class="font-medium">{{ record.weight }}kg</span>
                        </div>
                        <div v-if="record.headCirc" class="flex items-center">
                          <span class="text-gray-500 dark:text-gray-400 mr-2">头围：</span>
                          <span class="font-medium">{{ record.headCirc }}cm</span>
                        </div>
                      </div>
                      <p v-if="record.notes" class="mt-2 text-gray-600 dark:text-gray-300">
                        {{ record.notes }}
                      </p>
                    </div>
                    <UButton 
                      @click="editGrowthRecord(record)"
                      color="gray"
                      variant="ghost"
                      size="sm"
                      icon="i-heroicons-pencil-square"
                    />
                  </div>

                  <!-- 里程碑 -->
                  <div v-else-if="record.type === 'milestone'" class="flex items-start">
                    <div class="w-8 h-8 bg-purple-500 rounded-full flex items-center justify-center text-white text-sm mr-4 mt-1">
                      <UIcon name="i-heroicons-star" class="w-4 h-4" />
                    </div>
                    <div class="flex-1">
                      <div class="flex items-center mb-2">
                        <h4 class="font-medium text-gray-900 dark:text-white">
                          {{ record.title }}
                        </h4>
                        <UBadge 
                          v-if="record.isImportant"
                          color="red"
                          variant="soft"
                          size="xs"
                          class="ml-2"
                        >
                          重要
                        </UBadge>
                      </div>
                      <p v-if="record.description" class="text-gray-600 dark:text-gray-300 mb-2">
                        {{ record.description }}
                      </p>
                      <UBadge 
                        v-if="record.category"
                        color="blue"
                        variant="soft"
                        size="xs"
                      >
                        {{ record.category }}
                      </UBadge>
                    </div>
                    <UButton 
                      @click="editMilestone(record)"
                      color="gray"
                      variant="ghost"
                      size="sm"
                      icon="i-heroicons-pencil-square"
                    />
                  </div>

                  <!-- Memo记录 -->
                  <div v-else-if="record.type === 'memo'" class="flex items-start">
                    <div class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white text-sm mr-4 mt-1">
                      <UIcon name="i-heroicons-camera" class="w-4 h-4" />
                    </div>
                    <div class="flex-1">
                      <h4 class="font-medium text-gray-900 dark:text-white mb-2">成长记录</h4>
                      <div class="text-gray-600 dark:text-gray-300 mb-2" v-html="record.content"></div>
                      <div v-if="record.imgs" class="grid grid-cols-2 md:grid-cols-4 gap-2">
                        <img 
                          v-for="(img, index) in record.imgs.split(',')" 
                          :key="index"
                          :src="img"
                          :alt="`照片${index + 1}`"
                          class="w-full h-24 object-cover rounded"
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 添加成长记录模态框 -->
    <UModal v-model="showGrowthModal" :ui="{ width: 'sm:max-w-md' }">
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold">添加成长记录</h3>
            <UButton
              color="gray"
              variant="ghost"
              icon="i-heroicons-x-mark-20-solid"
              class="-my-1"
              @click="closeGrowthModal"
            />
          </div>
        </template>

        <form @submit.prevent="saveGrowthRecord" class="space-y-4">
          <UFormGroup label="记录日期" required>
            <UInput
              v-model="growthForm.recordDate"
              type="date"
              placeholder="请选择记录日期"
            />
          </UFormGroup>

          <UFormGroup label="身高 (cm)">
            <UInput
              v-model="growthForm.height"
              type="number"
              step="0.1"
              placeholder="请输入身高"
            />
          </UFormGroup>

          <UFormGroup label="体重 (kg)">
            <UInput
              v-model="growthForm.weight"
              type="number"
              step="0.1"
              placeholder="请输入体重"
            />
          </UFormGroup>

          <UFormGroup label="头围 (cm)">
            <UInput
              v-model="growthForm.headCirc"
              type="number"
              step="0.1"
              placeholder="请输入头围"
            />
          </UFormGroup>

          <UFormGroup label="备注">
            <UTextarea
              v-model="growthForm.notes"
              placeholder="请输入备注信息"
              rows="3"
            />
          </UFormGroup>
        </form>

        <template #footer>
          <div class="flex justify-end space-x-2">
            <UButton
              color="gray"
              variant="soft"
              @click="closeGrowthModal"
            >
              取消
            </UButton>
            <UButton
              color="green"
              @click="saveGrowthRecord"
              :loading="savingGrowth"
            >
              保存
            </UButton>
          </div>
        </template>
      </UCard>
    </UModal>

    <!-- 添加里程碑模态框 -->
    <UModal v-model="showMilestoneModal" :ui="{ width: 'sm:max-w-md' }">
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold">添加里程碑</h3>
            <UButton
              color="gray"
              variant="ghost"
              icon="i-heroicons-x-mark-20-solid"
              class="-my-1"
              @click="closeMilestoneModal"
            />
          </div>
        </template>

        <form @submit.prevent="saveMilestone" class="space-y-4">
          <UFormGroup label="标题" required>
            <UInput v-model="milestoneForm.title" placeholder="请输入里程碑标题" />
          </UFormGroup>

          <UFormGroup label="日期" required>
            <UInput
              v-model="milestoneForm.date"
              type="date"
              placeholder="请选择日期"
            />
          </UFormGroup>

          <UFormGroup label="分类">
            <USelect
              v-model="milestoneForm.category"
              :options="[
                { label: '运动', value: '运动' },
                { label: '语言', value: '语言' },
                { label: '认知', value: '认知' },
                { label: '社交', value: '社交' },
                { label: '其他', value: '其他' }
              ]"
              placeholder="请选择分类"
            />
          </UFormGroup>

          <UFormGroup label="描述">
            <UTextarea
              v-model="milestoneForm.description"
              placeholder="请输入描述信息"
              rows="3"
            />
          </UFormGroup>

          <UFormGroup>
            <UCheckbox
              v-model="milestoneForm.isImportant"
              label="标记为重要里程碑"
            />
          </UFormGroup>
        </form>

        <template #footer>
          <div class="flex justify-end space-x-2">
            <UButton
              color="gray"
              variant="soft"
              @click="closeMilestoneModal"
            >
              取消
            </UButton>
            <UButton
              color="purple"
              @click="saveMilestone"
              :loading="savingMilestone"
            >
              保存
            </UButton>
          </div>
        </template>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { ChildVO, GrowthRecordVO, MilestoneVO, MemoVO, UserVO } from '~/types'

const route = useRoute()
const currentUser = useState<UserVO>('userinfo')

const child = ref<ChildVO | null>(null)
const growthRecords = ref<GrowthRecordVO[]>([])
const milestones = ref<MilestoneVO[]>([])
const memos = ref<MemoVO[]>([])

const showGrowthModal = ref(false)
const showMilestoneModal = ref(false)
const savingGrowth = ref(false)
const savingMilestone = ref(false)

const selectedYear = ref('')
const selectedMonth = ref('')

const growthForm = reactive({
  childId: 0,
  height: '',
  weight: '',
  headCirc: '',
  recordDate: '',
  notes: ''
})

const milestoneForm = reactive({
  childId: 0,
  title: '',
  description: '',
  category: '',
  date: '',
  isImportant: false
})

// 年份选项
const yearOptions = computed(() => {
  const currentYear = new Date().getFullYear()
  const years = []
  for (let i = currentYear; i >= currentYear - 10; i--) {
    years.push({ label: `${i}年`, value: i.toString() })
  }
  return years
})

// 月份选项
const monthOptions = computed(() => {
  const months = []
  for (let i = 1; i <= 12; i++) {
    months.push({ label: `${i}月`, value: i.toString().padStart(2, '0') })
  }
  return months
})

// 合并所有记录并按日期分组
const groupedRecords = computed(() => {
  const allRecords: any[] = []
  
  // 添加成长记录
  growthRecords.value.forEach(record => {
    allRecords.push({
      ...record,
      type: 'growth',
      date: record.recordDate
    })
  })
  
  // 添加里程碑
  milestones.value.forEach(milestone => {
    allRecords.push({
      ...milestone,
      type: 'milestone',
      date: milestone.date
    })
  })
  
  // 添加Memo记录
  memos.value.forEach(memo => {
    allRecords.push({
      ...memo,
      type: 'memo',
      date: memo.createdAt
    })
  })
  
  // 按日期分组
  const grouped: Record<string, any[]> = {}
  allRecords
    .filter(record => {
      if (selectedYear.value && !record.date.startsWith(selectedYear.value)) return false
      if (selectedMonth.value && !record.date.includes(`-${selectedMonth.value}-`)) return false
      return true
    })
    .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime())
    .forEach(record => {
      const date = record.date.split('T')[0]
      if (!grouped[date]) {
        grouped[date] = []
      }
      grouped[date].push(record)
    })
  
  return grouped
})

// 获取儿童信息
const loadChild = async () => {
  try {
    const childId = route.params.id
    const res = await useMyFetch<ChildVO>(`/child/get/${childId}`)
    child.value = res
  } catch (error) {
    console.error('获取儿童信息失败:', error)
  }
}

// 获取成长记录
const loadGrowthRecords = async () => {
  try {
    const childId = route.params.id
    const res = await useMyFetch<GrowthRecordVO[]>(`/growth/list?childId=${childId}`)
    growthRecords.value = res
  } catch (error) {
    console.error('获取成长记录失败:', error)
  }
}

// 获取里程碑
const loadMilestones = async () => {
  try {
    const childId = route.params.id
    const res = await useMyFetch<MilestoneVO[]>(`/milestone/list?childId=${childId}`)
    milestones.value = res
  } catch (error) {
    console.error('获取里程碑失败:', error)
  }
}

// 获取Memo记录
const loadMemos = async () => {
  try {
    const childId = route.params.id
    // 这里需要根据实际的API调整
    const res = await useMyFetch<{ list: MemoVO[] }>('/memo/list', { childId })
    memos.value = res.list
  } catch (error) {
    console.error('获取Memo记录失败:', error)
  }
}

// 计算年龄
const calculateAge = (birthDate: string) => {
  const birth = new Date(birthDate)
  const today = new Date()
  let age = today.getFullYear() - birth.getFullYear()
  const monthDiff = today.getMonth() - birth.getMonth()
  
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
    age--
  }
  
  return age
}

// 格式化日期
const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('zh-CN')
}

// 格式化日期标题
const formatDateTitle = (date: string) => {
  const d = new Date(date)
  return d.toLocaleDateString('zh-CN', { month: 'long', day: 'numeric' })
}

// 格式化完整日期
const formatDateFull = (date: string) => {
  const d = new Date(date)
  return d.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric', weekday: 'long' })
}

// 格式化日期（仅显示日）
const formatDay = (date: string) => {
  const d = new Date(date)
  return d.getDate()
}

// 清除筛选
const clearFilter = () => {
  selectedYear.value = ''
  selectedMonth.value = ''
}

// 保存成长记录
const saveGrowthRecord = async () => {
  if (!growthForm.recordDate) {
    alert('请选择记录日期')
    return
  }
  
  savingGrowth.value = true
  try {
    growthForm.childId = Number(route.params.id)
    await useMyFetch('/growth/create', growthForm)
    await loadGrowthRecords()
    closeGrowthModal()
  } catch (error) {
    console.error('保存成长记录失败:', error)
  } finally {
    savingGrowth.value = false
  }
}

// 保存里程碑
const saveMilestone = async () => {
  if (!milestoneForm.title || !milestoneForm.date) {
    alert('请填写必填项')
    return
  }
  
  savingMilestone.value = true
  try {
    milestoneForm.childId = Number(route.params.id)
    await useMyFetch('/milestone/create', milestoneForm)
    await loadMilestones()
    closeMilestoneModal()
  } catch (error) {
    console.error('保存里程碑失败:', error)
  } finally {
    savingMilestone.value = false
  }
}

// 关闭成长记录模态框
const closeGrowthModal = () => {
  showGrowthModal.value = false
  Object.assign(growthForm, {
    childId: 0,
    height: '',
    weight: '',
    headCirc: '',
    recordDate: '',
    notes: ''
  })
}

// 关闭里程碑模态框
const closeMilestoneModal = () => {
  showMilestoneModal.value = false
  Object.assign(milestoneForm, {
    childId: 0,
    title: '',
    description: '',
    category: '',
    date: '',
    isImportant: false
  })
}

// 编辑成长记录
const editGrowthRecord = (record: GrowthRecordVO) => {
  // 实现编辑功能
  console.log('编辑成长记录:', record)
}

// 编辑里程碑
const editMilestone = (milestone: MilestoneVO) => {
  // 实现编辑功能
  console.log('编辑里程碑:', milestone)
}

onMounted(() => {
  loadChild()
  loadGrowthRecords()
  loadMilestones()
  loadMemos()
})
</script>