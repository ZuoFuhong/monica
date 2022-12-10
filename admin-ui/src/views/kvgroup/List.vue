<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">配置中心</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="服务名">
            <el-input v-model="formInline.name" placeholder="请输入服务名" maxlength="30"></el-input>
          </el-form-item>
          <el-form-item label="业务">
            <el-input v-model="formInline.business" placeholder="请输入业务名" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" @click="handleSearch">查询</el-button>
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
            <template slot-scope="scope">
              <el-button @click="doJumpEdit(scope.row)" type="text" size="small">{{ scope.row.name }}</el-button>
            </template>
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="namespace"
            label="环境名称">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="business"
            label="业务名称">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="op_user"
            label="最后发布人">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="ptime"
            label="最后发布时间">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
              <el-button @click="handleHistory(scope.row)" type="text" size="small">更改历史</el-button>
              <el-button @click="handleHistoryVersion(scope.row)" type="text" size="small">历史版本</el-button>
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
  name: "KvGroup",
  created() {
    console.log("init kvgroup page")
    this.serviceList = []
  },
  data() {
    return {
      formInline: {
        name: '',
        business: ''
      },
      serviceList: [
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a3',
          name: 'springboot.monica.portal',
          namespace: 'Production',
          business: '测试应用',
          op_user: 'mars',
          ctime: 1655716935,
          ptime: 1655716949,
          metadata: {}
        },
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a2',
          name: 'springboot.monica.portal.http',
          namespace: 'Test',
          business: '测试应用2',
          op_user: 'mars',
          ctime: 1655716935,
          ptime: 1655716949,
          metadata: {}
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
      console.log(this.formInline)
    },
    // 跳转编辑
    doJumpEdit(row) {
      this.$router.push(`/kvgroup/${row.namespace}/${row.name}`)
    },
    // 更改历史
    handleHistory(row) {
      console.log(row)
    },
    // 历史版本
    handleHistoryVersion(row) {
      console.log(row)
    },
    // 查看服务
    handleEdit(row) {
      this.$router.push(`/kvgroup/${row.namespace}/${row.name}`)
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