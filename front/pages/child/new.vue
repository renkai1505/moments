<template>
  <div class="min-h-screen bg-gray-50 dark:bg-neutral-900">
    <!-- 头部导航 -->
    <div class="bg-white dark:bg-neutral-800 shadow-sm">
      <div class="max-w-2xl mx-auto px-4 py-3">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-3">
            <UButton 
              icon="i-heroicons-arrow-left" 
              variant="ghost" 
              @click="navigateTo('/')"
            />
            <h1 class="text-xl font-bold">创建儿童档案</h1>
          </div>
        </div>
      </div>
    </div>

    <!-- 表单内容 -->
    <div class="max-w-2xl mx-auto px-4 py-6">
      <div class="bg-white dark:bg-neutral-800 rounded-lg shadow-lg p-6">
        <UForm :schema="schema" :state="form" @submit="handleSubmit" class="space-y-6">
          <!-- 基本信息 -->
          <div>
            <h2 class="text-lg font-semibold mb-4 flex items-center">
              <UIcon name="i-heroicons-user" class="w-5 h-5 mr-2" />
              基本信息
            </h2>
            
            <!-- 头像上传 -->
            <div class="flex items-center space-x-4 mb-6">
              <div class="relative">
                <UAvatar 
                  :src="form.avatarUrl || '/default-avatar.png'" 
                  :alt="form.name || '头像'" 
                  size="2xl"
                />
                <UButton 
                  icon="i-heroicons-camera" 
                  size="sm"
                  class="absolute -bottom-1 -right-1"
                  @click="uploadAvatar"
                />
              </div>
              <div>
                <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">点击相机图标上传头像</p>
                <p class="text-xs text-gray-500">建议尺寸: 200x200 像素</p>
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <UFormGroup label="姓名" name="name" required>
                <UInput v-model="form.name" placeholder="请输入宝贝的姓名" />
              </UFormGroup>

              <UFormGroup label="昵称" name="nickname">
                <UInput v-model="form.nickname" placeholder="请输入昵称（可选）" />
              </UFormGroup>

              <UFormGroup label="性别" name="gender" required>
                <USelectMenu 
                  v-model="form.gender" 
                  :options="genderOptions"
                  placeholder="请选择性别"
                />
              </UFormGroup>

              <UFormGroup label="出生日期" name="birthDate" required>
                <UInput 
                  v-model="form.birthDate" 
                  type="date" 
                  :max="new Date().toISOString().split('T')[0]"
                />
              </UFormGroup>

              <UFormGroup label="血型" name="bloodType">
                <USelectMenu 
                  v-model="form.bloodType" 
                  :options="bloodTypeOptions"
                  placeholder="请选择血型（可选）"
                />
              </UFormGroup>

              <div></div> <!-- 占位 -->
            </div>
          </div>

          <!-- 身体信息 -->
          <div>
            <h2 class="text-lg font-semibold mb-4 flex items-center">
              <UIcon name="i-heroicons-scale" class="w-5 h-5 mr-2" />
              身体信息
            </h2>
            
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <UFormGroup label="身高(cm)" name="height">
                <UInput 
                  v-model="form.height" 
                  type="number" 
                  placeholder="请输入身高"
                  min="0"
                  max="300"
                  step="0.1"
                />
              </UFormGroup>

              <UFormGroup label="体重(kg)" name="weight">
                <UInput 
                  v-model="form.weight" 
                  type="number" 
                  placeholder="请输入体重"
                  min="0"
                  max="200"
                  step="0.1"
                />
              </UFormGroup>
            </div>
          </div>

          <!-- 其他信息 -->
          <div>
            <h2 class="text-lg font-semibold mb-4 flex items-center">
              <UIcon name="i-heroicons-heart" class="w-5 h-5 mr-2" />
              其他信息
            </h2>
            
            <div class="space-y-4">
              <UFormGroup label="兴趣爱好" name="hobbies">
                <UTextarea 
                  v-model="form.hobbies" 
                  placeholder="请输入宝贝的兴趣爱好"
                  rows="3"
                />
              </UFormGroup>

              <UFormGroup label="简介描述" name="description">
                <UTextarea 
                  v-model="form.description" 
                  placeholder="请输入对宝贝的简介描述"
                  rows="4"
                />
              </UFormGroup>
            </div>
          </div>

          <!-- 封面图上传 -->
          <div>
            <h2 class="text-lg font-semibold mb-4 flex items-center">
              <UIcon name="i-heroicons-photo" class="w-5 h-5 mr-2" />
              封面图片
            </h2>
            
            <div class="border-2 border-dashed border-gray-300 dark:border-neutral-600 rounded-lg p-6 text-center">
              <div v-if="form.coverUrl" class="relative inline-block">
                <img :src="form.coverUrl" alt="封面图" class="max-w-xs max-h-48 rounded-lg">
                <UButton 
                  icon="i-heroicons-x-mark"
                  size="sm"
                  color="red"
                  class="absolute -top-2 -right-2"
                  @click="form.coverUrl = ''"
                />
              </div>
              <div v-else>
                <UIcon name="i-heroicons-photo" class="w-12 h-12 text-gray-400 mx-auto mb-4" />
                <p class="text-gray-600 dark:text-gray-400 mb-2">点击上传封面图片</p>
                <p class="text-xs text-gray-500">建议尺寸: 1200x400 像素</p>
                <UButton 
                  color="primary" 
                  variant="outline"
                  class="mt-4"
                  @click="uploadCover"
                >
                  选择图片
                </UButton>
              </div>
            </div>
          </div>

          <!-- 提交按钮 -->
          <div class="flex items-center justify-end space-x-4 pt-6 border-t border-gray-200 dark:border-neutral-700">
            <UButton 
              variant="ghost" 
              @click="navigateTo('/')"
            >
              取消
            </UButton>
            <UButton 
              type="submit" 
              color="primary"
              :loading="submitting"
            >
              创建档案
            </UButton>
          </div>
        </UForm>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { z } from 'zod';
import type { SaveChildReq } from "~/types";

// 表单验证schema
const schema = z.object({
  name: z.string().min(1, '请输入姓名'),
  nickname: z.string().optional(),
  gender: z.enum(['M', 'F'], { required_error: '请选择性别' }),
  birthDate: z.string().min(1, '请选择出生日期'),
  bloodType: z.string().optional(),
  height: z.number().min(0).optional(),
  weight: z.number().min(0).optional(),
  hobbies: z.string().optional(),
  description: z.string().optional(),
  avatarUrl: z.string().optional(),
  coverUrl: z.string().optional(),
});

// 表单数据
const form = ref<SaveChildReq>({
  name: '',
  nickname: '',
  gender: undefined,
  birthDate: '',
  bloodType: '',
  height: undefined,
  weight: undefined,
  hobbies: '',
  description: '',
  avatarUrl: '',
  coverUrl: '',
});

const submitting = ref(false);

// 选项数据
const genderOptions = [
  { label: '👦 男孩', value: 'M' },
  { label: '👧 女孩', value: 'F' }
];

const bloodTypeOptions = [
  { label: 'A型', value: 'A' },
  { label: 'B型', value: 'B' },
  { label: 'AB型', value: 'AB' },
  { label: 'O型', value: 'O' },
  { label: '未知', value: '未知' }
];

// 提交表单
const handleSubmit = async (data: SaveChildReq) => {
  try {
    submitting.value = true;
    
    const submitData = {
      ...data,
      // 确保数字类型正确
      height: data.height ? Number(data.height) : undefined,
      weight: data.weight ? Number(data.weight) : undefined,
    };

    const result = await useMyFetch<number>('/child/save', {
      method: 'POST',
      body: submitData
    });

    if (result) {
      // 成功创建，跳转到儿童详情页
      await navigateTo(`/child/${result}`);
    }
  } catch (error) {
    console.error('创建儿童档案失败:', error);
    // TODO: 显示错误提示
  } finally {
    submitting.value = false;
  }
};

// 上传头像
const uploadAvatar = async () => {
  // TODO: 实现文件上传
  console.log('上传头像');
};

// 上传封面
const uploadCover = async () => {
  // TODO: 实现文件上传
  console.log('上传封面');
};

// SEO
useSeoMeta({
  title: '创建儿童档案 - 儿童相册',
  description: '为您的宝贝创建专属的成长档案，记录美好时光'
});
</script>

<style scoped>
.max-w-xs {
  max-width: 20rem;
}

.max-h-48 {
  max-height: 12rem;
}
</style>