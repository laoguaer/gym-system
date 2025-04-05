import { baseRequest, request } from '@/utils/http'

export default {
  login: (data = {}) => baseRequest.post('/login', data),
  register: (data = {}) => baseRequest.post('/register', data),
  logout: () => baseRequest.get('/logout'),
  /** 发送验证码 */
  sendCode: params => baseRequest.get('/code', { params }),

  /** 关于我 */
  about: () => request.get('/about'),
  /** 获取页面 */
  getPageList: () => request.get('/page'),
  /** 首页数据 */
  getHomeData: () => request.get('/home'),
  /** 首页文章列表 */
  getArticles: params => request.get('/article/list', { params }),
  /** 文章详情 */
  getArticleDetail: id => request.get(`/article/${id}`),
  /** 文章归档 */
  getArchives: (params = {}) => request.get('/article/archive', { params }),
  /** 文章搜索 */
  searchArticles: (params = {}) => request.get('/article/search', { params }),
  /** 获取视频列表 */
  getVideos: params => request.get('/video/list', { params }),
  /** 获取视频列表 */
  getVideoDetail: id => request.get(`/video/${id}`),

  /** 菜单列表 */
  getCategorys: () => request.get('/category/list'),
  /** 标签列表 */
  getTags: () => request.get('/tag/list'),
  /** 留言列表 */
  getMessages: () => request.get('/message/list'),
  /** 友链列表 */
  getLinks: () => request.get('/link/list'),
  /** 评论列表 */
  getComments: (params = {}) => request.get('/comment/list', { params }),
  /** 评论回复列表 */
  getCommentReplies: (id, params = {}) => request.get(`/comment/replies/${id}`, { params }),

  // ! 需要 Token 的接口
  /** 根据 token 获取当前用户信息 */
  getUser: () => request.get('/user/info', { needToken: true }),
  /** 修改当前用户信息 */
  updateUser: data => request.put('/user/info', data, { needToken: true }),
  /** 留言 */
  saveMessage: data => request.post('/message', data, { needToken: true }),
  /** 评论 */
  saveComment: data => request.post('/comment', data, { needToken: true }),
  /** 点赞评论 */
  saveLikeComment: id => request.get(`/comment/like/${id}`, { needToken: true }),
  /** 点赞文章 */
  saveLikeArticle: id => request.get(`/article/like/${id}`, { needToken: true }),

  /** ChatBot 对话 */
  chatWithBot: async (data) => {
    try {
      const response = await request.post('chat/send', data)
      return response.data // 确保返回的是 response.data
    }
    catch (error) {
      // 处理错误，确保返回有意义的信息
      throw error.response ? error.response.data : error
    }
  },

  //* * 获取教练信息 */
  getTrainerInfo: (params = {}) => request.get('/coach/list', { params }),

  //* * 获取课程列表 */
  getCourseList: (params = {}) => request.get('/course/list', { params }),
  getCourseTags: () => request.get('/course/tags'),
  getMyCourseList: (params = {}) => request.get('/course/myCourse', { params }),
}
