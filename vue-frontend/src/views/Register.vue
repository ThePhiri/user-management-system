<template>
  <div class="container">
    <div class="row">
      <div class="col-md-4 order-md-2 mb-4"></div>
      <div class="col-md-8 order-md-1">
        <h4 class="mb-3">Registeration Form</h4>
        <form v-on:submit.prevent="submitForm">
          <div class="row">
            <div class="col-md-6 mb-3">
              <label for="firstName">First name</label>
              <input
                type="text"
                class="form-control"
                id="firstName"
                v-model="firstName"
                placeholder=""
                required
              />
            </div>
            <div class="col-md-6 mb-3">
              <label for="lastName">Last name</label>
              <input
                type="text"
                class="form-control"
                id="lastName"
                v-model="lastName"
                placeholder=""
                required
              />
            </div>
          </div>

          <div class="mb-3">
            <label for="email">Email </label>
            <input
              type="email"
              class="form-control"
              id="email"
              v-model="email"
              placeholder="you@example.com"
              required
            />
          </div>

          <div class="mb-3">
            <label for="phone">Phone</label>
            <input
              type="text"
              class="form-control"
              id="phone"
              v-model="phone"
              placeholder="0123456789"
              required
            />
          </div>

          <div class="mb-3">
            <label for="password">Password </label>
            <input
              type="text"
              class="form-control"
              id="password"
              v-model="password"
              placeholder="Password"
              required
            />

            <div
              v-if="password.length > 1 && password.length < 6"
              class="text-danger"
            >
              Password should be minimum 6 characters long
            </div>
          </div>
          <div class="mb-3">
            <label for="password2">Confirm Password </label>
            <input
              type="text"
              class="form-control"
              id="password2"
              v-model="password2"
              placeholder="Password"
              required
            />
            <div v-if="password != password2" class="text-danger">
              Passwords dont match
            </div>
          </div>

          <div class="row">
            <div class="col-md-5 mb-3">
              <label for="userType">User Type</label>
              <select
                class="custom-select d-block w-100"
                id="userType"
                v-model="userType"
                required
              >
                <option value="">Choose...</option>
                <option>ADMIN</option>
                <option>USER</option>
              </select>
            </div>
          </div>

          <button
            @click="register"
            class="btn btn-primary btn-lg btn-block"
            type="submit"
          >
            Register
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      firstName: "",
      lastName: "",
      email: "",
      phone: "",
      password: "",
      password2: "",
      userType: "",
    };
  },
  methods: {
    register: async function () {
      var myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");

      var raw = JSON.stringify({
        first_name: this.firstName,
        last_name: this.lastName,
        password: this.password,
        email: this.email,
        phone: this.phone,
        user_type: this.userType,
      });

      var requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow",
      };

      fetch("http://localhost:8080/users/signup", requestOptions)
        .then((response) => response.text())
        .then((result) => console.log(result))
        .catch((error) => console.log("error", error));
    },
  },
};
</script>

<style scoped>
</style>