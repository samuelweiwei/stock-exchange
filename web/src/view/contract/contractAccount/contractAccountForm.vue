
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联用户表的用户 ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总保证金:" prop="totalMargin">
          <el-input-number v-model="formData.totalMargin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="可用保证金:" prop="availableMargin">
          <el-input-number v-model="formData.availableMargin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="冻结保证金:" prop="frozenMargin">
          <el-input-number v-model="formData.frozenMargin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="已用保证金:" prop="usedMargin">
          <el-input-number v-model="formData.usedMargin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="已实现盈亏:" prop="realizedProfitLoss">
          <el-input-number v-model="formData.realizedProfitLoss" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="委托状态:" prop="STATUS">
          <el-input v-model.number="formData.STATUS" :clearable="true" placeholder="请输入" />
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
  createContractAccount,
  updateContractAccount,
  findContractAccount
} from '@/api/contract/contractAccount'

defineOptions({
    name: 'ContractAccountForm'
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
            userId: undefined,
            totalMargin: 0,
            availableMargin: 0,
            frozenMargin: 0,
            usedMargin: 0,
            realizedProfitLoss: 0,
            STATUS: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findContractAccount({ ID: route.query.id })
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
               res = await createContractAccount(formData.value)
               break
             case 'update':
               res = await updateContractAccount(formData.value)
               break
             default:
               res = await createContractAccount(formData.value)
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
