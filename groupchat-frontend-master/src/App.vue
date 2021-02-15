<template>
  <div
    class="enter"
    v-if="!roomOpen"
  >
    <Auth
      @token="token = $event"
      v-if="token === ''"
    />
    <button
      v-else
      @click="token = ''"
    >로그아웃</button>
    <hr>
    <RoomSelection
      :roomID="selectedRoom ? selectedRoom.id : 0"
      @selectChange="selectedRoom = $event"
    />
    <hr>
    <button
      class="enterroombutton"
      @click="roomOpen = true"
      :disabled="token === ''"
    >방 들어가기</button>
  </div>

  <Messenger
    v-else
    :token="token"
    :roomID="selectedRoom.id"
    :roomName="selectedRoom.name"
    @disconnect="roomOpen = false"
  />
</template>

<script>
import { ref } from 'vue'
import Messenger from "./components/Messenger.vue"
import RoomSelection from './components/RoomSelection.vue'
import Auth from './components/Auth.vue'

export default {
  name: "App",
  components: {
    Messenger,
    Auth,
    RoomSelection
  },
  setup() {
    const selectedRoom = ref(undefined)
    const token = ref('')
    const roomOpen = ref(false)
    return {
      selectedRoom,
      token,
      roomOpen
    }
  }
};
</script>

<style>
body {
  margin: 0;
  padding: 0;
  background-color: #bbc9e0;
}

.enter {
  width: 200px;
  margin: auto;
  background: #fff;
  padding: 15px;
  border-radius: 5px;
  box-shadow: rgba(0, 0, 0, 0.1) 0 0 10px;
  margin-top: 100px;
}

.enterroombutton {
  box-sizing: border-box;
  width: 100%;
}
</style>
