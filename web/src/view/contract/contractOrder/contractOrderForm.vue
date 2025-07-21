
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="orderNumber">
          <el-input v-model="formData.orderNumber" :clearable="true"  placeholder="请输入订单号" />
       </el-form-item>
        <el-form-item label="订单时间:" prop="orderTime">
          <el-date-picker v-model="formData.orderTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="关联用户表的用户 ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="股票id:" prop="stockId">
          <el-input v-model.number="formData.stockId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="股票名称:" prop="stockName">
          <el-input v-model="formData.stockName" :clearable="true"  placeholder="请输入股票名称" />
       </el-form-item>
        <el-form-item label="下单类型:" prop="orderType">
          <el-input v-model.number="formData.orderType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开仓价格:" prop="openPrice">
          <el-input-number v-model="formData.openPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="平仓价格:" prop="closePrice">
          <el-input-number v-model="formData.closePrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="操作类型:" prop="operationType">
          <el-input v-model.number="formData.operationType" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input-number v-model="formData.quantity" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="杠杆倍数:" prop="leverageRatio">
          <el-input v-model.number="formData.leverageRatio" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单状态:" prop="STATUS">
          <el-input v-model.number="formData.STATUS" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手续费:" prop="fee">
          <el-input-number v-model="formData.fee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="已实现盈亏:" prop="realizedProfitLoss">
          <el-input-number v-model="formData.realizedProfitLoss" :precision="2" :clearable="true"></el-input-number>
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
  createContractOrder,
  updateContractOrder,
  findContractOrder
} from '@/api/contract/contractOrder'

defineOptions({
    name: 'ContractOrderForm'
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
            orderNumber: '',
            orderTime: new Date(),
            userId: undefined,
            stockId: undefined,
            stockName: '',
            orderType: undefined,
            openPrice: 0,
            closePrice: 0,
            operationType: undefined,
            quantity: 0,
            leverageRatio: undefined,
            STATUS: undefined,
            fee: 0,
            realizedProfitLoss: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findContractOrder({ ID: route.query.id })
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
               res = await createContractOrder(formData.value)
               break
             case 'update':
               res = await updateContractOrder(formData.value)
               break
             default:
               res = await createContractOrder(formData.value)
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
