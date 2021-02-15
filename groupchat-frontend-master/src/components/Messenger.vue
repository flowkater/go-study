<template>
  <MenuBar :roomName="props.roomName" @disconnect="disconnect" />
  <div class="messages">
    <Message
      v-for="(message, idx) in messages"
      :message="message"
      :mine="myID === message.senderID"
      :displaySender="(myID !== message.senderID) && (messages[idx-1] ?? {sender: null}).senderID !== message.senderID"
      :key="idx"
    />
  </div>

  <div class="message-input-bar">
    <MessageInput @send="onSend" />
  </div>
</template>

<script>
import { ref } from "vue";
import Message from "./Message.vue";
import MessageInput from "./MessageInput.vue";
import MenuBar from "./MenuBar.vue";
export default {
  components: { Message, MessageInput, MenuBar },
  name: "Messenger",
  emits: ['disconnect'],
  props: {
    token: String,
    roomID: Number,
    roomName: String
  },
  setup(props, {emit}) {
    const myID = ref('');
    const socket = new WebSocket(`ws://localhost:8080/room/${props.roomID}?token=${encodeURIComponent(props.token)}`);
    const messages = ref([]);

    const parseJwt = (token) => {
      try {
        return JSON.parse(atob(token.split('.')[1]));
      } catch (e) {
        return null;
      }
    };

    const parsedToken = parseJwt(props.token)
    myID.value = parsedToken.uid

    const onSend = message =>
      socket.send(
        JSON.stringify({
          text: message,
        })
      );
    socket.onmessage = messageEvent =>
      messages.value.push(JSON.parse(messageEvent.data));
    return {
      myID,
      messages,
      onSend,
      props,
      disconnect: () => emit('disconnect')
    };
  }
};
</script>

<style scoped>
.messages {
  width: 500px;
  height: 100%;
  margin: auto;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: stretch;

  flex-grow: 1;
  overflow: auto;
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: stretch;

  padding-top: 70px;
  padding-bottom: 70px;
}

.message-input-bar {
  height: 60px;
  width: 100%;
  position: fixed;
  bottom: 0;
  background-color: #bbc9e0;
}
</style>