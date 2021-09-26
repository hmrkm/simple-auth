<template>
  <div id="login">
    <div class="h-100 bg-plum-plate-login bg-animation">
      <div class="d-flex h-100 justify-content-center align-items-center">
        <b-col md="8" class="mx-auto app-login-box">
          <div class="modal-dialog w-100 mx-auto">
            <div class="modal-content">
              <div class="modal-body">
                <form ref="form">
                  <b-form-group id="emailGroup1" label-for="emailLabel1">
                    <b-form-input
                      id="emailInput"
                      type="text"
                      name="email"
                      v-model="email"
                      required
                      placeholder="Enter email..."
                    >
                    </b-form-input>
                  </b-form-group>
                  <b-form-group
                    id="exampleInputGroup2"
                    label-for="exampleInput2"
                  >
                    <b-form-input
                      id="exampleInput2"
                      type="password"
                      name="password"
                      v-model="password"
                      required
                      placeholder="Enter password..."
                    >
                    </b-form-input>
                  </b-form-group>
                </form>
                <div class="divider" />
              </div>
              <div class="modal-footer clearfix">
                <div class="float-right">
                  <b-button variant="primary" @click="doLogin" size="lg"
                    >Login</b-button
                  >
                </div>
              </div>
            </div>
          </div>
        </b-col>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "Login",
  data: () => ({
    email: "",
    password: "",
  }),
  methods: {
    ...mapActions("login", {
      login: "login",
    }),
    reset() {
      this.$refs.form.reset();
    },
    doLogin() {
      const { email, password } = this;
      if (this.email != "" && this.password != "") {
        this.login({ email, password }).then((result) => {
          if (result) {
            this.$store.commit(`message/setContent`, {
              content: "認証に成功しました",
              timeout: 3000,
              clss: "success",
            });
          } else {
            this.$store.commit(`message/setContent`, {
              content: "認証に失敗しました",
              timeout: 3000,
              clss: "danger",
            });
          }
        });
      } else {
        alert("Please fill the text!");
      }
    },
  },
};
</script>
