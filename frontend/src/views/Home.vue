<template>
  <div>
    <Message />
    <Login />
  </div>
</template>

<script>
import Message from "@/components/Message";
import Login from "@/components/Login";

export default {
  name: "Home",
  components: {
    Message,
    Login,
  },
};
</script>
