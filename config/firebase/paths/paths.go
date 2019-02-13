package paths

import (
	"strings"
)

func ExtractAndReplaceMethodIdAndDatePath(
	id string,
	networkPath string,
) string {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1] -> {Root}/{Uid}/{Date}/{Snowflake}.
	//Extract everything but snowflake.
	splits := strings.Split(strings.Split(networkPath, "/documents/")[1], "/")[:3]
	//Split by collection and extract [1] which should be fromId.
	fromId := splits[1]

	dbPath := strings.Replace(strings.Join(splits, "/"), fromId, id, 1)

	return dbPath
}

func ExtractAndReplaceMethodIdAndDatePathWithSnowflake(
	id string,
	networkPath string,
) (string, string) {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1] -> {Root}/{Uid}/{Date}/{Snowflake}.
	path := strings.Split(networkPath, "/documents/")[1]
	//Split by collection and extract [1] which should be fromId.
	pathSplits := strings.Split(path, "/")
	fromId := pathSplits[1]
	//Extract [3] which is document unique id.
	docId := pathSplits[3]
	//Extract everything but snowflake.
	dbPath := strings.Replace(strings.Join(pathSplits[:3], "/"), fromId, id, 1)

	return docId, dbPath
}

func ExtractMethodIdAndDatePath(
	networkPath string,
) string {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1] -> {Root}/{Uid}/{Date}/{Snowflake}.
	path := strings.Split(networkPath, "/documents/")[1]
	//Split by collection and extract [1] which should be fromId.
	splits := strings.Split(path, "/")[:3]

	//Extract everything but snowflake.
	dbPath := strings.Join(splits, "/")

	return dbPath
}

func ExtractMethodIdAndDatePathWithSnowflake(
	networkPath string,
) (string, string) {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1] -> {Root}/{Uid}/{Date}/{Snowflake}.
	path := strings.Split(networkPath, "/documents/")[1]
	splits := strings.Split(path, "/")

	dbPath := strings.Join(splits[:3], "/")

	return splits[3], dbPath
}

func TransformGroupIntoMonetary(
	groupPath string,
) string {
	return strings.Replace(groupPath, "GroupRequests", "MonetaryRequests", 1)
}
