
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="工厂ID:" prop="factoryId">
          <el-input v-model.number="formData.factoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父分类id:" prop="parentId">
          <el-input v-model.number="formData.parentId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="祖级列表:" prop="ancestors">
          <el-input v-model="formData.ancestors" :clearable="true"  placeholder="请输入祖级列表" />
       </el-form-item>
        <el-form-item label="安全教育名称:" prop="categoryName">
          <el-input v-model="formData.categoryName" :clearable="true"  placeholder="请输入安全教育名称" />
       </el-form-item>
        <el-form-item label="用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
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
  createSafetyEducationCategories,
  updateSafetyEducationCategories,
  findSafetyEducationCategories
} from '@/api/stock/safetyEducationCategories'

defineOptions({
    name: 'SafetyEducationCategoriesForm'
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
            factoryId: undefined,
            parentId: undefined,
            ancestors: '',
            categoryName: '',
            userId: undefined,
        })
// 验证规则
const rule = reactive({
               factoryId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSafetyEducationCategories({ ID: route.query.id })
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
               res = await createSafetyEducationCategories(formData.value)
               break
             case 'update':
               res = await updateSafetyEducationCategories(formData.value)
               break
             default:
               res = await createSafetyEducationCategories(formData.value)
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
