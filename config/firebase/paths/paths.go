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

	//Split at gcpstuff in [0] and full db path in [1].
	path := strings.Split(networkPath, "/documents/")[1]
	//Split by collection and extract [1] which should be fromId.
	fromId := strings.Split(path, "/")[1]
	dbPath := strings.Replace(path, fromId, id, 1)

	return dbPath
}

func ExtractAndReplaceMethodIdAndDatePathWithSnowflake(
	id string,
	networkPath string,
) (string, string) {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1].
	path := strings.Split(networkPath, "/documents/")[1]
	//Split by collection and extract [1] which should be fromId.
	pathSplits := strings.Split(path, "/")
	fromId := pathSplits[1]
	//Extract [3] which is document unique id.
	docId := pathSplits[3]
	dbPath := strings.Replace(path, fromId, id, 1)

	return docId, dbPath
}

func ExtractMethodIdAndDatePath(
	networkPath string,
) string {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1].
	return strings.Split(networkPath, "/documents/")[1]
}

func ExtractMethodIdAndDatePathWithSnowflake(
	networkPath string,
) (string, string) {

	// Use the debtor id to split the path
	// so it gets the path next to the creditor ID
	// to be replaced by the debtor ID, so it inserts on the correct collection.

	//Split at gcpstuff in [0] and full db path in [1].
	path := strings.Split(networkPath, "/documents/")[1]
	splits := strings.Split(path, "/")
	return splits[3], path
}

func TransformGroupIntoMonetary(
	networkPath string,
) string {
	return strings.Replace(networkPath, "GroupRequests", "MonetaryRequests", 1)
}
