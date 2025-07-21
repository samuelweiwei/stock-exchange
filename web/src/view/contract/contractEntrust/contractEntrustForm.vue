
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关联用户表的用户 ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联持仓表的持仓 ID:" prop="positionId">
          <el-input v-model.number="formData.positionId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联订单表的订单 ID:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="股票id:" prop="stockId">
          <el-input v-model.number="formData.stockId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="股票名称:" prop="stockName">
          <el-input v-model="formData.stockName" :clearable="true"  placeholder="请输入股票名称" />
       </el-form-item>
        <el-form-item label="触发类型:" prop="triggerType">
          <el-input v-model.number="formData.triggerType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="触发价格:" prop="triggerPrice">
          <el-input-number v-model="formData.triggerPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="操作类型:" prop="operationType">
          <el-input v-model.number="formData.operationType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input-number v-model="formData.quantity" :precision="2" :clearable="true"></el-input-number>
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
  createContractEntrust,
  updateContractEntrust,
  findContractEntrust
} from '@/api/contract/contractEntrust'

defineOptions({
    name: 'ContractEntrustForm'
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
            positionId: undefined,
            orderId: undefined,
            stockId: undefined,
            stockName: '',
            triggerType: undefined,
            triggerPrice: 0,
            operationType: undefined,
            quantity: 0,
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
      const res = await findContractEntrust({ ID: route.query.id })
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
               res = await createContractEntrust(formData.value)
               break
             case 'update':
               res = await updateContractEntrust(formData.value)
               break
             default:
               res = await createContractEntrust(formData.value)
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
