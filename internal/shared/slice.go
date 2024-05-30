package shared

func ChunkSlice[T any](items []T, chunkSize int) [][]T {
	if len(items) == 0 {
		return [][]T{}
	}

	if chunkSize < 1 {
		return [][]T{
			items,
		}
	}

	chunks := make([][]T, 0, len(items)/chunkSize)

	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
