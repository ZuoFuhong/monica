import {get, post, _delete} from "../utils/axios"
import {Base64} from 'js-base64'

class Service {

  /**
   * 保存服务
   */
  async saveService(data) {
    return await post('/api/v1/service', data)
  }

  /**
   * 搜索服务
   */
  async searchService(data) {
    let bval = Base64.encode(JSON.stringify(data))
    let query = encodeURIComponent(bval)
    return await get(`/api/v1/services?query=${query}`)
  }

  /**
   * 删除单个服务
   */
  async deleteService(name) {
    return await _delete(`/api/v1/service?name=${name}`)
  }

  /**
   * 查询服务命名空间
   */
  async searchServiceNS(data) {
    let bval = Base64.encode(JSON.stringify(data))
    let query = encodeURIComponent(bval)
    return await get(`/api/v1/service_namespaces?query=${query}`)
  }

  /**
   * 搜索服务实例
   */
  async searchServiceInstance(data) {
    let bval = Base64.encode(JSON.stringify(data))
    let query = encodeURIComponent(bval)
    return await get(`/api/v1/service_instances?query=${query}`)
  }

  /**
   * 保存服务实例
   */
  async saveServiceInstance(data) {
    return await post('/api/v1/service_instance', data)
  }

  /**
   * 删除服务实例
   */
  async deleteServiceInstance(ns, sname, ip) {
    return await _delete(`/api/v1/service_instance?ns=${ns}&sname=${sname}&ip=${ip}`)
  }
}

export default new Service()
