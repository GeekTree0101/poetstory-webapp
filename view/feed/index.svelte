<script>
  export let feeds = []
</script>

<h1>Feed Index</h1>

<table border="1" cellpadding="10">
  {#if feeds.length > 0}
    <thead>
      {#each Object.keys(feeds[0]) as key}
        <th>{key}</th>
      {/each}
    </thead>
  {/if}
  {#each feeds as feed}
    <tr>
      {#each Object.keys(feed) as key}
        {#if key.toLowerCase() === "id"}
          <td><a href={`/feed/${feed.id || 0}`}>{feed[key]}</a></td>
        {:else}
          <td>{feed[key]}</td>
        {/if}
      {/each}
    </tr>
  {/each}
</table>

<br />

<a href={`/feed/new`}>New Feed</a>

<style>
  table {
    border-collapse: collapse;
  }
</style>
