<template>
  <div class="roomselection">
    <select
      :selected="props.roomID"
      @change="selectChange"
      v-if="rooms.length > 0"
    >
      <option
        :value="room.id"
        v-for="(room, idx) in rooms"
        :key="idx"
      >{{room.name}}</option>
    </select>

    <form @submit.prevent="addRoom">
      <input
        class="roomnameinput"
        type="text"
        placeholder="방 이름"
        v-model="roomName"
      />
      <button class="plusbutton">+</button>
    </form>

  </div>
</template>

<script>
import { ref } from 'vue'
export default {
  name: 'RoomSelection',
  props: {
    roomID: Number
  },
  emit: ['selectChange'],
  setup(props, { emit }) {
    const rooms = ref([])
    const roomName = ref('')

    const fetchRooms = async () => {
      const resp = await fetch('http://localhost:8080/rooms')
      const parsed = await resp.json()
      rooms.value = parsed
    }

    fetchRooms().then(() => emit('selectChange', rooms.value[0]))

    const addRoom = async () => {
      await fetch('http://localhost:8080/room', {
        method: 'POST',
        body: JSON.stringify({
          'name': roomName.value
        })
      })
      if (rooms.value.length === 0) {
          fetchRooms().then(() => emit('selectChange', rooms.value[0]))
      } else {
          fetchRooms()
      }
      roomName.value = ''
    }

    return {
      props,
      rooms,
      roomName,
      addRoom,
      selectChange: (e) => emit('selectChange', rooms.value.find(room => room.id == e.target.value))
    }
  }
}
</script>

<style scoped>
.roomselection {
  font-family: sans-serif;
  width: 100%;
  margin: auto;
}

select {
  box-sizing: border-box;
  width: 100%;
  padding: 2px 5px;
}

form {
  box-sizing: border-box;
  width: 100%;
  display: flex;
  margin: 5px 0;
}

.roomnameinput {
  flex-grow: 1;
  padding: 2px 5px;
}

.plusbutton {
  padding: 2px 5px;
  flex-grow: 0;
}
</style>