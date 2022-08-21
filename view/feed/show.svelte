<script>
  export let feed = {}
</script>

<h1>Show Feed</h1>

<table border="1" cellpadding="10">
  <thead>
    {#each Object.keys(feed) as key}
      <th>{key}</th>
    {/each}
  </thead>
  <tr>
    {#each Object.keys(feed) as key}
      <td>{feed[key]}</td>
    {/each}
  </tr>
</table>

<br />

<a href={`/feed`}>Back</a>
<span> | </span>
<a href={`/feed/${feed.id || 0}/edit`}>Edit</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
