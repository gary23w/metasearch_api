package wikipedia

import (
	"context"
	"net/url"
	"testing"

	"github.com/gary23w/metasearch_api/internal/search"
	"github.com/gary23w/metasearch_api/search/internal/searchtest"
	"github.com/stretchr/testify/require"
)

func mustURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return *u
}

func TestWiki(t *testing.T) {
	s := New()
	ctx := context.Background()
	const qu = "Solar System"
	resp, err := s.SearchRaw(ctx, SearchReq{
		Titles: qu,
		Prop: []Property{
			PropExtracts,
			PropPageImages,
		},
	})
	const extract = `The Solar System is the gravitationally bound system of the Sun and the objects that orbit it, either directly or indirectly, including the eight planets and five dwarf planets as defined by the International Astronomical Union (IAU). Of the objects that orbit the Sun directly, the largest eight are the planets, with the remainder being smaller objects, such as dwarf planets and small Solar System bodies. Of the objects that orbit the Sun indirectly—the moons—two are larger than the smallest planet, Mercury.The Solar System formed 4.6 billion years ago from the gravitational collapse of a giant interstellar molecular cloud. The vast majority of the system's mass is in the Sun, with the majority of the remaining mass contained in Jupiter. The four smaller inner planets, Mercury, Venus, Earth and Mars, are terrestrial planets, being primarily composed of rock and metal. The four outer planets are giant planets, being substantially more massive than the terrestrials. The two largest, Jupiter and Saturn, are gas giants, being composed mainly of hydrogen and helium; the two outermost planets, Uranus and Neptune, are ice giants, being composed mostly of substances with relatively high melting points compared with hydrogen and helium, called volatiles, such as water, ammonia and methane. All eight planets have almost circular orbits that lie within a nearly flat disc called the ecliptic.
The Solar System also contains smaller objects. The asteroid belt, which lies between the orbits of Mars and Jupiter, mostly contains objects composed, like the terrestrial planets, of rock and metal. Beyond Neptune's orbit lie the Kuiper belt and scattered disc, which are populations of trans-Neptunian objects composed mostly of ices, and beyond them a newly discovered population of sednoids. Within these populations are several dozen to possibly tens of thousands of objects large enough that they have been rounded by their own gravity. Such objects are categorized as dwarf planets. Identified dwarf planets include the asteroid Ceres and the trans-Neptunian objects Pluto and Eris. In addition to these two regions, various other small-body populations, including comets, centaurs and interplanetary dust clouds, freely travel between regions. Six of the planets, at least four of the dwarf planets, and many of the smaller bodies are orbited by natural satellites, usually termed "moons" after the Moon. Each of the outer planets is encircled by planetary rings of dust and other small objects.
The solar wind, a stream of charged particles flowing outwards from the Sun, creates a bubble-like region in the interstellar medium known as the heliosphere. The heliopause is the point at which pressure from the solar wind is equal to the opposing pressure of the interstellar medium; it extends out to the edge of the scattered disc. The Oort cloud, which is thought to be the source for long-period comets, may also exist at a distance roughly a thousand times further than the heliosphere. The Solar System is located in the Orion Arm, 26,000 light-years from the center of the Milky Way galaxy.`
	require.NoError(t, err)
	require.Equal(t, SearchResp{
		BatchComplete: true,
		Query: QueryResponse{
			Pages: []Page{
				{
					ID:      26903,
					Title:   "Solar System",
					Extract: extract,
					Thumbnail: &Image{
						Source: "https://upload.wikimedia.org/wikipedia/commons/thumb/c/cb/Planets2013.svg/300px-Planets2013.svg.png",
						Width:  300,
						Height: 177,
					},
					PageImage: "Planets2013.svg",
				},
			},
		},
	}, *resp)
	t.Logf("%#v", resp)

	it := s.Search(ctx, search.Request{
		Query: qu,
	})
	defer it.Close()

	var got []search.Result
	for it.Next(ctx) {
		got = append(got, it.Result())
	}
	require.NoError(t, err)
	require.Equal(t, []search.Result{
		&search.EntityResult{
			LinkResult: search.LinkResult{
				URL:   mustURL("https://en.wikipedia.org/wiki/Solar_System"),
				Title: "Solar System",
				Desc:  extract,
			},
			Image: &search.Image{
				URL:    mustURL("https://upload.wikimedia.org/wikipedia/commons/thumb/c/cb/Planets2013.svg/300px-Planets2013.svg.png"),
				Width:  300,
				Height: 177,
			},
		},
	}, got)
}

func TestWikipedia(t *testing.T) {
	s := New()
	searchtest.RunSearchTest(t, s, nil)
}
