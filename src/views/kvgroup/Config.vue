<template>
  <div>
    <el-row class="content-header">
      <p class="page-title">配置项 {{namespace}}/{{service_name}}</p>
    </el-row>
    <div class="content-body">
      <el-row>
        <el-form :inline="true" :model="formInline">
          <el-form-item>
            <el-input v-model="formInline.kw" placeholder="请输入key过滤" maxlength="30" prefix-icon="el-icon-search"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button icon="el-icon-plus" @click="handleAdd">新增配置</el-button>
            <el-button icon="el-icon-s-promotion" type="success" @click="handlePublish">发布</el-button>
            <el-button icon="el-icon-back" type="info">回滚</el-button>
          </el-form-item>
        </el-form>
      </el-row>
      <el-row>
        <el-table
          :data="configItems"
          border
          style="width: 100%">
          <el-table-column
            align="center"
            prop="step"
            label="发布状态"
            width="100">
            <template slot-scope="scope">
              <span v-if="scope.row.step === 0" class="add-step">新增</span>
              <span v-if="scope.row.step === 1" class="processing-step">发布中</span>
              <span v-if="scope.row.step === 2" class="release-step">已发布</span>
              <span v-if="scope.row.step === 3" class="edit-step">有修改</span>
              <span v-if="scope.row.step === 4" class="del-step">已删除</span>
            </template>
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="key"
            label="Key">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="value"
            label="Value">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="remark"
            label="备注">
          </el-table-column>
          <el-table-column
            show-overflow-tooltip
            prop="op_user"
            label="最后修改人">
          </el-table-column>
          <el-table-column
            :formatter="formatTime"
            show-overflow-tooltip
            prop="mtime"
            label="最后修改时间">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
              <el-button type="text" size="small">编辑</el-button>
              <el-button type="text" size="small" v-if="scope.row.step === 3 || scope.row.step === 4">撤销</el-button>
              <el-button type="text" size="small">删除</el-button>
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
  name: 'KvGroupConfig',
  created() {
    this.namespace = this.$route.params.ns
    this.service_name = this.$route.params.name
  },
  data() {
    return {
      namespace: "",
      service_name: "",
      formInline: {
        kw: '',
        kwFlag: 'key'
      },
      // 配置项：0-新增 1-发布中 2-已发布 3-有修改 4-已删除
      configItems: [
        {
          id: 'abcaka39dafjj23lkjda03k4y',
          step: 0,
          key: 'key1',
          value: '12345',
          remark: 'test',
          op_user: 'mars',
          mtime: 1655716935
        },
        {
          id: 'abcaka39dafjj23lkjda03k4y',
          step: 1,
          key: 'key2',
          value: '12345',
          remark: 'test',
          op_user: 'mars',
          mtime: 1655716935
        },
        {
          id: 'abcaka39dafjj23lkjda03k4y',
          step: 2,
          key: 'key3',
          value: '12345',
          remark: 'test',
          op_user: 'mars',
          mtime: 1655716935
        },
        {
          id: 'abcaka39dafjj23lkjda03k4y',
          step: 3,
          key: 'key4',
          value: '12345',
          remark: 'test',
          op_user: 'mars',
          mtime: 1655716935
        },
        {
          id: 'abcaka39dafjj23lkjda03k4y',
          step: 4,
          key: 'key5',
          value: '12345',
          remark: 'test',
          op_user: 'mars',
          mtime: 1655716935
        }
      ],
      total: 100,
      pageSize: 10,
      currentPage: 1
    }
  },
  methods: {
    // 切换分页
    handleCurrentChange(val) {
      // Todo: 分页加载
      console.log(val)
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
.add-step{
  font-size: 12px;
  color: #ffffff;
  background-color: #67c23a;
  padding: 4px 14px;
  font-weight: bold;
}
.release-step {
  font-size: 12px;
  color: #ffffff;
  background-color: #909399;
  padding: 4px 8px;
  font-weight: bold;
}
.processing-step {
  font-size: 12px;
  color: #ffffff;
  background-color: #f0ad4e;
  padding: 4px 8px;
  font-weight: bold;
}
.del-step {
  font-size: 12px;
  color: #ffffff;
  background-color: #f56c6c;
  padding: 4px 8px;
  font-weight: bold;
}
.edit-step {
  font-size: 12px;
  color: #ffffff;
  background-color: #409eff;
  padding: 4px 8px;
  font-weight: bold;
}
</style>