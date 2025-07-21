
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="userId字段:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="1 充值 2 提现 3 划转到合约账户 4 合约账户划转进来:" prop="transactionType">
          <el-input v-model="formData.transactionType" :clearable="true"  placeholder="请输入1 充值 2 提现 3 划转到合约账户 4 合约账户划转进来" />
       </el-form-item>
        <el-form-item label="amount字段:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="balanceAfter字段:" prop="balanceAfter">
          <el-input-number v-model="formData.balanceAfter" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="transactionDate字段:" prop="transactionDate">
          <el-date-picker v-model="formData.transactionDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="description字段:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入description字段" />
       </el-form-item>
        <el-form-item label="订单id:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true"  placeholder="请输入订单id" />
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
  createUserAccountFlow,
  updateUserAccountFlow,
  findUserAccountFlow
} from '@/api/userfund/userAccountFlow'

defineOptions({
    name: 'UserAccountFlowForm'
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
            transactionType: '',
            amount: 0,
            balanceAfter: 0,
            transactionDate: new Date(),
            description: '',
            orderId: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserAccountFlow({ ID: route.query.id })
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
               res = await createUserAccountFlow(formData.value)
               break
             case 'update':
               res = await updateUserAccountFlow(formData.value)
               break
             default:
               res = await createUserAccountFlow(formData.value)
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
