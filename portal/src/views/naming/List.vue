<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">服务治理</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item label="服务名">
            <el-input v-model="formInline.name" placeholder="请输入服务名" maxlength="30"></el-input>
          </el-form-item>
          <el-form-item label="实例IP">
            <el-input v-model="formInline.host" placeholder="请输入实例IP" maxlength="20"></el-input>
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
              <el-button @click="doJumpInstance(scope.row)" type="text" size="small">{{ scope.row.name }}</el-button>
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
            align="center"
            prop="instance_count"
            label="实例数"
            width="150">
            <template slot-scope="scope">
              <el-button @click="doJumpInstance(scope.row)" type="text" size="small">{{ scope.row.instance_count }}</el-button>
            </template>
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
export default {
  name: "NamingService",
  created() {
    console.log("init service page")
  },
  data() {
    return {
      formInline: {
        name: '',
        host: '',
        business: ''
      },
      serviceList: [
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a3',
          name: 'springboot.monica.portal',
          namespace: 'Production',
          business: '测试应用',
          instance_count: 1,
          ctime: 1655716935,
          mtime: 1655716949
        },
        {
          id: 'd990964439b442b8b7adb5e2f15ac6a2',
          name: 'springboot.monica.portal.http',
          namespace: 'Test',
          business: '测试应用2',
          instance_count: 5,
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
      console.log(this.formInline)
    },
    // 查看服务实例
    doJumpInstance(row) {
      this.$router.push(`/naming-service/instance/${row.namespace}/${row.name}`)
    },
    // 查看服务Token
    handlePreview(row) {
      // Todo: 加载服务Token
      console.log(row)
      this.$alert('d880964439b442b8b7adb5e2f15ac6a2', '服务 Token', {
        confirmButtonText: '确定'
      });
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