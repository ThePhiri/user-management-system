<template>
  <div class="container">
    <div class="row">
      <h1>All users</h1>
    </div>
    <div class="row">
      <div class="col-md-3"></div>
      <div class="col-md-6">
        <ul
          v-for="user in users"
          :key="user.id"
          class="list-group list-group-flush"
        >
          <!-- <router-link :to="{ name: 'UsersDetails', params: { id: user.id } }"> -->
          <li class="list-group-item">{{ user.name }}</li>
          <!-- </router-link> -->
        </ul>
      </div>
      <div class="col-md-3"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: "allusers",
  data() {
    return {
      users: [],
    };
  },

  mounted() {
    var myHeaders = new Headers();
    myHeaders.append(
      "token",
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImVtYWlsQGRvbWFpbjIuY29tIiwiRmlyc3RfbmFtZSI6IlBoaXJpIiwiTGFzdF9uYW1lIjoiUGhpcmkiLCJVaWQiOiI2MTFhNDRkNmFjOGMwMWEzMzY0OTc5YTYiLCJVc2VyX3R5cGUiOiJBRE1JTiIsImV4cCI6MTYyOTMwODMzOH0.qqDpXnsDT2xStfoFUEzSc9zcro_9pobew0RyT-GaeAM"
    );

    var requestOptions = {
      method: "GET",
      headers: myHeaders,
      redirect: "follow",
    };

    fetch("http://localhost:8080/users", requestOptions)
      .then((response) => response.json())
      .then((data) => (this.allusers = data))
      .catch((error) => console.log("error", error));
  },

  // methods: {
  //   allusers: async function () {
  //     var myHeaders = new Headers();
  //     myHeaders.append(
  //       "token",
  //       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImVtYWlsQGRvbWFpbjIuY29tIiwiRmlyc3RfbmFtZSI6IlBoaXJpIiwiTGFzdF9uYW1lIjoiUGhpcmkiLCJVaWQiOiI2MTFhNDRkNmFjOGMwMWEzMzY0OTc5YTYiLCJVc2VyX3R5cGUiOiJBRE1JTiIsImV4cCI6MTYyOTMwODMzOH0.qqDpXnsDT2xStfoFUEzSc9zcro_9pobew0RyT-GaeAM"
  //     );

  //     var requestOptions = {
  //       method: "GET",
  //       headers: myHeaders,
  //       redirect: "follow",
  //     };

  //     fetch("http://localhost:8080/users", requestOptions)
  //       .then((response) => response.json())
  //       .then((data) => (this.allusers = data))
  //       .catch((error) => console.log("error", error));
  //   },
  // },
};
</script>

<style>
</style>