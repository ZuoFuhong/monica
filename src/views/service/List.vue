<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">服务列表</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="服务名">
            <el-input v-model="formInline.serviceName" placeholder="请输入服务名" maxlength="30"></el-input>
          </el-form-item>
          <el-form-item label="业务">
            <el-input v-model="formInline.businessName" placeholder="请输入业务名" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" @click="handleSearch">查询</el-button>
            <el-button icon="el-icon-plus" @click="handleAdd">新建</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row>
        <el-table
          :data="serviceList"
          border
          style="width: 100%">
          <el-table-column
            show-overflow-tooltip
            prop="name"
            label="服务名">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="business"
            label="业务名称">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="owners"
            label="负责人">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="remark"
            label="备注">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="ctime"
            label="创建时间">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="mtime"
            label="修改时间">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="handleEdit(scope.row)" type="text" size="small">编辑</el-button>
              <el-button @click="handleDelete(scope.row)" type="text" size="small"><span class="deleteBtn">删除</span></el-button>
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
    </div>
  </div>
</template>

<script>
import moment from 'moment';
export default {
  name: "AppList",
  created() {
    console.log("init service page")
  },
  data() {
    return {
      formInline: {
        serviceName: '',
        businessName: ''
      },
      serviceList: [
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a3',
          name: 'test',
          business: '测试应用',
          owners: 'mars',
          remark: "123",
          ctime: 1655716935,
          mtime: 1655716949
        },
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a2',
          name: 'test',
          business: '测试应用2',
          owners: 'mars',
          remark: "123",
          ctime: 1655716935,
          mtime: 1655716949
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    // 搜索服务
    handleSearch() {
      console.log('submit!');
    },
    // 新建服务
    handleAdd() {
    },
    // 编辑服务
    handleEdit(row) {
      console.log(row)
    },
    // 删除服务
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
.deleteBtn{
  color: red;
}
.deleteBtn:hover{
  color: rgb(251, 100, 100);
}
</style>