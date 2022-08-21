<script>
  export let poet = {}
</script>

<h1>Show Poet</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(poet) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(poet) as key}
      <td>{poet[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href={`/poet`}>Back</a>
<span> | </span>
<a href={`/poet/${poet.id || 0}/edit`}>Edit</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
