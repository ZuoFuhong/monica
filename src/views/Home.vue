<template>
  <el-container>
    <el-main>
      <el-row style="text-align: left;">
        <el-form :inline="true" :model="formInline" class="demo-form-inline">
          <el-form-item label="服务名">
            <el-input v-model="formInline.serviceName" placeholder="请输入服务名" maxlength="30"></el-input>
          </el-form-item>
          <el-form-item label="实例IP">
            <el-input v-model="formInline.instance" placeholder="请输入实例IP" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item label="业务">
            <el-input v-model="formInline.businessName" placeholder="请输入业务名" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" @click="onSubmit">查询</el-button>
            <el-button icon="el-icon-plus" @click="onSubmit">新增</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row>
        <el-table
          :data="serviceList"
          border
          style="width: 100%">
          <el-table-column
            prop="name"
            label="服务名"
            width="180">
          </el-table-column>
          <el-table-column
            prop="namespace"
            label="环境名称"
            width="180">
          </el-table-column>
          <el-table-column
            prop="business"
            label="业务名称">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            prop="ctime"
            label="创建时间">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            prop="mtime"
            label="修改时间">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.row)" type="text" size="small">编辑</el-button>
              <el-button @click="handleDelete(scope.row)" type="text" size="small">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-row>
      <el-row style="text-align: right; padding: 20px 0;">
        <el-pagination
          @current-change="handleCurrentChange"
          :current-page.sync="currentPage"
          :page-size="pageSize"
          layout="total, prev, pager, next"
          :total="total">
        </el-pagination>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
import moment from 'moment';
export default {
  name: 'Home',
  data() {
    return {
      formInline: {
        serviceName: '',
        instance: '',
        businessName: ''
      },
      serviceList: [
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a3',
          name: 'test',
          namespace: 'Test',
          business: '测试应用',
          ctime: 1655716935,
          mtime: 1655716949,
          metadata: {}
        },
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a2',
          name: 'test',
          namespace: 'Test',
          business: '测试应用2',
          ctime: 1655716935,
          mtime: 1655716949,
          metadata: {}
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    onSubmit() {
      console.log('submit!');
    },
    handleEdit(row) {
      console.log(row)
    },
    handleDelete(row) {
      console.log(row)
    },
    // 时间格式
    formatTime(row, column, cellValue) {
      // 转毫秒时间戳
      return moment(cellValue * 1000).format("YYYY-MM-DD HH:mm:ss");
    },
    // 切换分页
    handleCurrentChange(val) {
      // Todo: 分页加载
      console.log(val)
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
