
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="主键:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="归属用户ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券ID:" prop="couponId">
          <el-input v-model.number="formData.couponId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券名称:" prop="couponName">
          <el-input v-model="formData.couponName" :clearable="true"  placeholder="请输入优惠券名称" />
       </el-form-item>
        <el-form-item label="优惠券金额:" prop="couponAmount">
          <el-input-number v-model="formData.couponAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="有效期开始时间:" prop="validStart">
          <el-input v-model.number="formData.validStart" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="有效期开始时间:" prop="validEnd">
          <el-input v-model.number="formData.validEnd" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="使用时间:" prop="usedTime">
          <el-input v-model.number="formData.usedTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="删除时间:" prop="deletedAt">
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
  createCouponRecord,
  updateCouponRecord,
  findCouponRecord
} from '@/api/coupon/couponRecord'

defineOptions({
    name: 'CouponRecordForm'
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
            userId: undefined,
            couponId: undefined,
            couponName: '',
            couponAmount: 0,
            validStart: undefined,
            validEnd: undefined,
            usedTime: undefined,
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
      const res = await findCouponRecord({ ID: route.query.id })
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
               res = await createCouponRecord(formData.value)
               break
             case 'update':
               res = await updateCouponRecord(formData.value)
               break
             default:
               res = await createCouponRecord(formData.value)
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
