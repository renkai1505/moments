<template>
  <div class="min-h-screen bg-gray-50 dark:bg-gray-900">
    <Header v-bind:user="currentUser"/>
    
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">儿童档案</h1>
        <p class="text-gray-600 dark:text-gray-400">管理您孩子的成长档案</p>
      </div>

      <!-- 添加儿童按钮 -->
      <div class="mb-6">
        <UButton 
          @click="showAddModal = true"
          color="blue"
          icon="i-heroicons-plus"
        >
          添加儿童档案
        </UButton>
      </div>

      <!-- 儿童档案列表 -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div 
          v-for="child in children" 
          :key="child.id"
          class="bg-white dark:bg-gray-800 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200"
        >
          <div class="p-6">
            <div class="flex items-center mb-4">
              <UAvatar 
                :src="child.avatar || '/default-child-avatar.svg'" 
                :alt="child.name"
                size="lg"
                class="mr-4"
              />
              <div>
                <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
                  {{ child.name }}
                </h3>
                <p class="text-sm text-gray-500 dark:text-gray-400">
                  {{ child.gender === '男' ? '男孩' : '女孩' }}
                  <span v-if="child.birthDate" class="ml-2">
                    {{ calculateAge(child.birthDate) }}岁
                  </span>
                </p>
              </div>
            </div>

            <div class="space-y-2 mb-4">
              <div v-if="child.birthDate" class="text-sm text-gray-600 dark:text-gray-300">
                <span class="font-medium">出生日期：</span>
                {{ formatDate(child.birthDate) }}
              </div>
              <div v-if="child.bloodType" class="text-sm text-gray-600 dark:text-gray-300">
                <span class="font-medium">血型：</span>
                {{ child.bloodType }}
              </div>
              <div v-if="child.description" class="text-sm text-gray-600 dark:text-gray-300">
                <span class="font-medium">描述：</span>
                {{ child.description }}
              </div>
            </div>

            <div class="flex space-x-2">
              <UButton 
                @click="viewChild(child)"
                color="gray"
                variant="soft"
                size="sm"
              >
                查看详情
              </UButton>
              <UButton 
                @click="editChild(child)"
                color="blue"
                variant="soft"
                size="sm"
              >
                编辑
              </UButton>
              <UButton 
                @click="deleteChild(child)"
                color="red"
                variant="soft"
                size="sm"
              >
                删除
              </UButton>
            </div>
          </div>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-if="children.length === 0" class="text-center py-12">
        <div class="text-gray-400 dark:text-gray-500 mb-4">
          <UIcon name="i-heroicons-user-group" class="w-16 h-16 mx-auto" />
        </div>
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">
          还没有儿童档案
        </h3>
        <p class="text-gray-500 dark:text-gray-400 mb-4">
          点击上方按钮添加第一个儿童档案
        </p>
      </div>
    </div>

    <!-- 添加/编辑儿童档案模态框 -->
    <UModal v-model="showAddModal" :ui="{ width: 'sm:max-w-md' }">
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold">
              {{ editingChild ? '编辑儿童档案' : '添加儿童档案' }}
            </h3>
            <UButton
              color="gray"
              variant="ghost"
              icon="i-heroicons-x-mark-20-solid"
              class="-my-1"
              @click="closeModal"
            />
          </div>
        </template>

        <form @submit.prevent="saveChild" class="space-y-4">
          <UFormGroup label="姓名" required>
            <UInput v-model="childForm.name" placeholder="请输入姓名" />
          </UFormGroup>

          <UFormGroup label="性别" required>
            <USelect
              v-model="childForm.gender"
              :options="[
                { label: '男孩', value: '男' },
                { label: '女孩', value: '女' }
              ]"
              placeholder="请选择性别"
            />
          </UFormGroup>

          <UFormGroup label="出生日期">
            <UInput
              v-model="childForm.birthDate"
              type="date"
              placeholder="请选择出生日期"
            />
          </UFormGroup>

          <UFormGroup label="血型">
            <USelect
              v-model="childForm.bloodType"
              :options="[
                { label: 'A型', value: 'A' },
                { label: 'B型', value: 'B' },
                { label: 'AB型', value: 'AB' },
                { label: 'O型', value: 'O' }
              ]"
              placeholder="请选择血型"
            />
          </UFormGroup>

          <UFormGroup label="头像">
            <UInput v-model="childForm.avatar" placeholder="头像URL" />
          </UFormGroup>

          <UFormGroup label="描述">
            <UTextarea
              v-model="childForm.description"
              placeholder="请输入描述信息"
              rows="3"
            />
          </UFormGroup>
        </form>

        <template #footer>
          <div class="flex justify-end space-x-2">
            <UButton
              color="gray"
              variant="soft"
              @click="closeModal"
            >
              取消
            </UButton>
            <UButton
              color="blue"
              @click="saveChild"
              :loading="saving"
            >
              {{ editingChild ? '更新' : '创建' }}
            </UButton>
          </div>
        </template>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { ChildVO, UserVO } from '~/types'

const currentUser = useState<UserVO>('userinfo')
const children = ref<ChildVO[]>([])
const showAddModal = ref(false)
const editingChild = ref<ChildVO | null>(null)
const saving = ref(false)

const childForm = reactive({
  name: '',
  gender: '',
  birthDate: '',
  bloodType: '',
  avatar: '',
  description: ''
})

// 获取儿童档案列表
const loadChildren = async () => {
  try {
    const res = await useMyFetch<ChildVO[]>('/child/list')
    children.value = res
  } catch (error) {
    console.error('获取儿童档案列表失败:', error)
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

// 查看儿童详情
const viewChild = (child: ChildVO) => {
  navigateTo(`/children/${child.id}`)
}

// 编辑儿童档案
const editChild = (child: ChildVO) => {
  editingChild.value = child
  Object.assign(childForm, {
    name: child.name,
    gender: child.gender,
    birthDate: child.birthDate ? child.birthDate.split('T')[0] : '',
    bloodType: child.bloodType,
    avatar: child.avatar,
    description: child.description
  })
  showAddModal.value = true
}

// 删除儿童档案
const deleteChild = async (child: ChildVO) => {
  if (!confirm(`确定要删除 ${child.name} 的档案吗？`)) {
    return
  }
  
  try {
    await useMyFetch(`/child/delete/${child.id}`)
    await loadChildren()
  } catch (error) {
    console.error('删除儿童档案失败:', error)
  }
}

// 保存儿童档案
const saveChild = async () => {
  if (!childForm.name || !childForm.gender) {
    alert('请填写必填项')
    return
  }
  
  saving.value = true
  try {
    if (editingChild.value) {
      await useMyFetch(`/child/update/${editingChild.value.id}`, childForm)
    } else {
      await useMyFetch('/child/create', childForm)
    }
    
    await loadChildren()
    closeModal()
  } catch (error) {
    console.error('保存儿童档案失败:', error)
  } finally {
    saving.value = false
  }
}

// 关闭模态框
const closeModal = () => {
  showAddModal.value = false
  editingChild.value = null
  Object.assign(childForm, {
    name: '',
    gender: '',
    birthDate: '',
    bloodType: '',
    avatar: '',
    description: ''
  })
}

onMounted(() => {
  loadChildren()
})
</script>