<script lang="ts">
  let formData = {
    Username: '',
    Password: '',
    CharName: ''
  };

  async function addPlayer() {
    try {
      // TODO: Figure out how to make .env var be api url and it not error 
      const response = await fetch("http://127.0.0.1:8090/players", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData) // Ensure formData is correctly populated
      });

      if (!response.ok) {
        throw new Error(`Failed to add player: ${response.status} ${response.statusText}`);
      }

      console.log('Player added successfully!');
    } catch (error) {
      console.error('Error adding player:', error.message);
    }
  }
</script>

<h2>Add New Player</h2>

<form on:submit|preventDefault={addPlayer}>
  <label for="username">Username:</label>
  <input type="text" id="username" bind:value={formData.Username} required />

  <label for="password">Password:</label>
  <input type="password" id="password" bind:value={formData.Password} required />

  <label for="charname">Character Name:</label>
  <input type="text" id="charname" bind:value={formData.CharName} required />

  <button type="submit">Add Player</button>
</form>
