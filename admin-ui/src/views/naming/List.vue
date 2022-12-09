<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">服务治理</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="服务名">
            <el-input v-model="formInline.kw" placeholder="请输入服务名" maxlength="30"></el-input>
          </el-form-item>
          <el-form-item label="实例IP">
            <el-input v-model="formInline.ip" placeholder="请输入实例IP" maxlength="20"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-search" @click="handleSearch">查询</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row>
        <el-table
          :data="snsList"
          border
          style="width: 100%">
          <el-table-column
            show-overflow-tooltip
            prop="name"
            label="服务名">
            <template slot-scope="scope">
              <el-button @click="doJumpInstance(scope.row)" type="text" size="small">{{ scope.row.name }}</el-button>
            </template>
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="ns"
            label="环境名称">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="business"
            label="业务名称">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="ctime"
            label="创建时间">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            fixed="right"
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button @click="handlePreview(scope.row)" type="text" size="small">查看 Token</el-button>
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
import serv from '../../model/service'
export default {
  name: "NamingService",
  created() {
    this.handleSearch()
  },
  data() {
    return {
      formInline: {
        kw: '',
        ip: ''
      },
      snsList: [],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    // 搜索服务
    async handleSearch() {
      let {retcode, errmsg, data} = await serv.searchServiceNS({
        kw: this.formInline.kw,
        ip: this.formInline.ip,
        page: this.currentPage,
        page_size: this.pageSize
      })
      if (retcode !== 0) {
        this.$message.error(errmsg)
        return
      }
      this.snsList = data.list
      this.total = data.total
    },
    // 切换分页
    async handleCurrentChange(val) {
      this.currentPage = val
      this.handleSearch()
    },
    // 查看服务实例
    doJumpInstance(row) {
      this.$router.push(`/naming-service/instance/${row.ns}/${row.name}`)
    },
    // 查看服务Token
    handlePreview(row) {
      this.$alert(row.token, '服务 Token', {
        confirmButtonText: '确定'
      }).catch(e=>e)
    },
    // 时间格式
    formatTime(row, column, cellValue) {
      // 转毫秒时间戳
      return moment(cellValue * 1000).format("YYYY-MM-DD HH:mm:ss");
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