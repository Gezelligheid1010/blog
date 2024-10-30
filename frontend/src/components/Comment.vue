<template>
  <div class="comment-section">
    <!-- 评论输入区 -->
    <div class="comment-input">
      <el-avatar :src="user.avatarImageUrl" class="user-avatar" v-if="user"></el-avatar>
      <el-input
          type="textarea"
          v-model="newComment"
          placeholder="来说两句吧…"
          rows="3"
      ></el-input>
      <el-button type="primary" @click="submitComment">畅言一下</el-button>
    </div>

    <!-- 评论列表 -->
<!--    <div class="comment-list">-->
    <div class="comment-container">
      <h3>最新评论</h3>
      <div v-for="comment in comments" :key="comment.author_id" class="comment">
<!--      <div v-for="comment in comments" :key="comment.author_id" class="comment-item">-->
        <el-avatar :src="comment.author.avatar" class="comment-avatar"></el-avatar>
        <div class="comment-content">
          <p>{{ comment.author.username }}</p>
          <p>{{ comment.content }}</p>
          <small>{{ comment.create_time | formatDate }}</small>
          <el-button type="text" icon="el-icon-like" @click="likeComment(comment.comment_id)">
            {{ comment.likes }}
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
// import axios from 'axios';

export default {
  name: "Comment",
  props: {
    sourceId: {
      type: String,
      required: true, // 强制要求父组件传递 sourceId
    },
  },
  data() {
    return {
      newComment: '', // 新的评论内容
      comments: [],   // 评论列表
      user: null,     // 当前登录用户信息
    };
  },
  mounted() {
    this.loadComments(); // 在组件挂载时加载评论
  },
  // created() {
  //   this.loadComments();
  //   this.user = JSON.parse(localStorage.getItem('loginResult')); // 获取当前用户信息
  //   console.log("loginResult:",this.user)
  // },
  methods: {
    async loadComments() {
      console.log("this.sourceId:",this.sourceId)
      this.$axios({
        method: "get",
        url: "/comment",
        params: {
          post_id : this.sourceId
        }
      }).then(response => {
        this.comments = response.data
          console.log("loadComments:",response);
        })
        .catch(error => {
          console.error('加载评论失败:', error);
        });
    },
    async submitComment() {
      if (!this.newComment) return;

      this.$axios({
        method: "post",
        url: "/comment",
        data: {
          content: this.newComment,
          post_id : this.sourceId
        }
      }).then(response => {
          console.log("submitComment:",response);
          if (response.code == 1000) {
            this.newComment = '';
            this.loadComments(); // 重新加载评论
          } else {
            console.log(response.msg);
          }
        })
        .catch(error => {
          console.error('提交评论失败:', error);
        });
    },

    async likeComment(commentId) {

      this.$axios({
        method: "post",
        url: `/comment/${commentId}`,
      }).then(response => {
          console.log("likeComment:",response);
          // console.log("commentId:",commentId);
          this.loadComments();
        })
        .catch(error => {
          console.error('点赞失败:', error);
        });
    }
  },
  filters: {
    formatDate(value) {
      const date = new Date(value);
      return date.toLocaleString();
    }
  }
};
</script>
<style scoped>
.comment-section {
  padding: 20px;
}

.comment-input {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-avatar {
  width: 40px;
  height: 40px;
}

.comment-list {
  margin-top: 20px;
}

.comment-item {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.comment-avatar {
  width: 30px;
  height: 30px;
  margin-right: 10px;
}

.comment-content {
  flex: 1;
}
.comment-container {
  margin: 20px 0;
}
.comment {
  padding: 10px;
  border-bottom: 1px solid #eaeaea;
}
</style>
