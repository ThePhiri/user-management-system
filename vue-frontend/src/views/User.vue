<template>
  <div class="container">
    <div class="row">
      <h1>Users dets</h1>
    </div>
    <div class="row">
      <div class="col-md-3"></div>
      <div class="col-md-6">
        <ul
          v-for="user in allusers"
          :key="user.ID"
          class="list-group list-group-flush"
        >
          <router-link :to="{ name: 'Users', params: { ID: user.ID } }">
            <li class="list-group-item">{{ user.name }}</li>
          </router-link>
        </ul>
      </div>
      <div class="col-md-3"></div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      allusers: [
        { ID: "1", name: "Chipo", email: "chipo@chip", phone: "123" },
        { ID: "2", name: "Chipo", email: "chipo@chip", phone: "2123" },
        { ID: "3", name: "Chipo", email: "chipo@chip", phone: "3123" },
      ],
    };
  },
};
</script>

<style>
</style>