<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
  
    // Define a writable store to hold player data
    const players = writable([]);
  
    async function fetchPlayers() {
      try {
        const response = await fetch('http://127.0.0.1:8090/players');
        if (!response.ok) {
          throw new Error('Failed to fetch player data');
        }
        const data = await response.json();
        players.set(data.data);
      } catch (error) {
        console.error('Error fetching players:', error.message);
      }
    }
  </script>
  
  <div>
    <h2>Player List</h2>
    <button on:click={fetchPlayers}>Fetch Players</button>
  
    {#if $players.length > 0}
      <ul>
        {#each $players as player}
          <li>Username: {player.Username} - CharName: {player.CharName}</li>
        {/each}
      </ul>
    {:else}
      <p>No players found</p>
    {/if}
  </div>