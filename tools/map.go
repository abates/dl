package main

import (
	"context"
	"encoding/xml"
	"io/ioutil"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmapi"
)

func main() {
	doc := &osm.OSM{}

	buf, _ := ioutil.ReadFile("SummitPoint.xml")
	xml.Unmarshal(buf, doc)

	for _, way := range doc.Ways {
		nodeIDS := make([]osm.NodeID, len(way.Nodes))
		for i, node := range way.Nodes {
			nodeIDS[i] = node.ID
		}

		nodes, err := osmapi.Nodes(context.Background(), nodeIDS)
		if err == nil {
			for i, node := range nodes {
				way.Nodes[i].Version = node.Version
				way.Nodes[i].ChangesetID = node.ChangesetID
				way.Nodes[i].Lat = node.Lat
				way.Nodes[i].Lon = node.Lon
			}
		}
	}
}
