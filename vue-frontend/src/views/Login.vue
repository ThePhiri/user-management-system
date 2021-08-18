<template>
  <div class="container">
    <div class="row">
      <h1>Login to see users</h1>
    </div>
    <div class="row">
      <div class="col-md-3"></div>
      <div class="col-md-6">
        <div class="mb-3">
          <label for="userEmail" class="form-label">Email address</label>
          <input
            type="email"
            class="form-control"
            id="userEmail"
            v-model="userEmail"
            aria-describedby="emailHelp"
            required
          />
        </div>
        <div class="mb-3">
          <label for="userPassword" class="form-label">Password</label>
          <input
            type="password"
            class="form-control"
            id="userPassword"
            v-model="userPassword"
            required
          />
          <div
            v-if="userPassword.length > 0 && userPassword.length < 6"
            class="text-danger"
          >
            Password should be minimum 6 characters long
          </div>
        </div>

        <button @click="login" class="btn btn-primary">Login</button>
      </div>
      <div class="col-md-3"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      userEmail: "",
      userPassword: "",
    };
  },

  methods: {
    login: async function () {
      var myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");

      var raw = { email: this.userEmail, password: this.userPassword };

      var requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: JSON.stringify(raw),
        redirect: "follow",
      };

      await fetch("http://localhost:8080/users/login", requestOptions)
        .then((response) => response.text())
        .then((result) => console.log(result))
        .catch((error) => console.log("error", error));
    },
  },
};
</script>

<style scoped>
</style>