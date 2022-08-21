<script>
  export let poets = []
</script>

<h1>Poet Index</h1>

<table border="1" cellpadding="10">
  {#if poets.length > 0}
    <thead>
      {#each Object.keys(poets[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each poets as poet}
    <tr>
      {#each Object.keys(poet) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/poet/${poet.id || 0}`}>{poet[key]}</a></td>
        {:else}
          <td>{poet[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<br />

<a href={`/poet/new`}>New Poet</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
