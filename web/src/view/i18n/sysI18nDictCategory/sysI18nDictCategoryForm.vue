
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="分类ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父分类ID:" prop="parentCategoryId">
          <el-input v-model.number="formData.parentCategoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分类名称:" prop="categoryName">
          <el-input v-model="formData.categoryName" :clearable="true"  placeholder="请输入分类名称" />
       </el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="deletedAt字段:" prop="deletedAt">
          <el-date-picker v-model="formData.deletedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSysI18nDictCategory,
  updateSysI18nDictCategory,
  findSysI18nDictCategory
} from '@/api/i18n/sysI18nDictCategory'

defineOptions({
    name: 'SysI18nDictCategoryForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: undefined,
            parentCategoryId: undefined,
            categoryName: '',
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSysI18nDictCategory({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSysI18nDictCategory(formData.value)
               break
             case 'update':
               res = await updateSysI18nDictCategory(formData.value)
               break
             default:
               res = await createSysI18nDictCategory(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
