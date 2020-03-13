package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title, Artist, Album string
	Year                 int
	Length               time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	dur, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return dur
}

func printTracks(tracks []*Track) {
	//http://networkbit.ch/golang-column-print/
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	table := new(tabwriter.Writer)
	//Init(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint)
	table.Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(table, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(table, format, "-----", "-----", "-----", "-----", "-----")
	for _, track := range tracks {
		fmt.Fprintf(table, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	table.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Customer sort multi-tier, first : Title, if title match, sort by Year, if Year match, sort by length
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// Order my length
type byLength []*Track

func (track byLength) Len() int           { return len(track) }
func (track byLength) Less(i, j int) bool { return track[i].Length < track[j].Length }
func (track byLength) Swap(i, j int)      { track[i], track[j] = track[j], track[i] }

func main() {

	//my_tracks := byArtist(tracks)
	//sort.Sort(my_tracks)
	//printTracks(my_tracks)

	//sort.Sort(customSort{ tracks,
	//	func(x, y *Track) bool {
	//	if x.Title != y.Title {
	//		return x.Title < y.Title
	//	}
	//	if x.Year != y.Year {
	//		return x.Year < y.Year
	//	}
	//	if x.Length != y.Length {
	//		return x.Length < y.Length
	//	}
	//	return false
	//}})
	// customSort
	// printTracks(tracks)

	sort.Sort(byLength(tracks))
	printTracks(tracks)
	// Good example : https://gobyexample.com/sorting-by-functions
}
