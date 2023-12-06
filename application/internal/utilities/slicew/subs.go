package slicew

func Subs[N any](items []N, numberOfSubslices int) (subs [][]N) {
	var (
		normalBucketSize = len(items) / numberOfSubslices
		bigBucketSize    = normalBucketSize + 1
		overflow         = len(items) - normalBucketSize*numberOfSubslices
	)
	for bucketIndex := 0; bucketIndex < numberOfSubslices; bucketIndex++ {
		var (
			previousBigBuckets    = min(bucketIndex, overflow)
			previousNormalBuckets = bucketIndex - previousBigBuckets
			start                 = previousBigBuckets*bigBucketSize + previousNormalBuckets*normalBucketSize
		)
		var end int
		if bucketIndex < overflow {
			end = start + bigBucketSize
		} else {
			end = start + normalBucketSize
		}
		subs = append(subs, items[start:end])
	}
	return
}
