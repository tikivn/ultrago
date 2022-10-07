package u_slice

func Chunk[T any](slice []T, size int) [][]T {
	if size < 1 {
		panic("chunk size cannot be less than 1")
	}
	var chunks [][]T
	for i := 0; i < len(slice); {
		// Clamp the last chunk to the slice bound as necessary.
		end := size
		if l := len(slice[i:]); l < size {
			end = l
		}

		// Set the capacity of each chunk so that appending to a chunk does not
		// modify the original slice.
		chunks = append(chunks, slice[i:i+end:i+end])
		i += end
	}

	return chunks
}
