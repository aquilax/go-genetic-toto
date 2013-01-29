package main

import "flag"

type Options struct {
	help      bool
	version   bool
	pool_size int
	num_draws int

	draws string
}

func (o *Options) Init() *flag.FlagSet {
	var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
	fs.BoolVar(&(options.help), "help", false, "Shows this message")
	fs.BoolVar(&(options.version), "version", false, "Show program version")

	fs.StringVar(&(options.draws), "draws", "", "Draws file name")
	fs.IntVar(&(options.pool_size), "pool_size", 100, "Number of chromosomes in population")
	fs.IntVar(&(options.num_draws), "num_draws", 100, "Number of random draws to generate if no file is set")
	return fs
}
