<template>
  <div
    class="alert"
    v-bind:class="{ 'alert-success': isSuccess, 'alert-danger': isDanger }"
    v-show="message"
  >
    {{ message }}
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  computed: {
    ...mapState({
      message: (state) => state.message.content,
      isSuccess: (state) => state.message.clss === "success",
      isDanger: (state) => state.message.clss === "danger",
    }),
  },
};
</script>
