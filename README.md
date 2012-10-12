 ================================================================================
 OPTARGEX - A simple commandline options parser.
================================================================================

 - Allows options with single and multi character names according to the
   traditional unix way of doing things. eg:  -o versus --option
 - Allows multi character only named options, eg: --verbose without -v
 - Exposes a channel based iterator which returns parsed options from the
   optargex.Parse() function. Note that it only yields Options which are actually
   present in the commandline arguments (os.Args). Call as a for loop:

   for opt := range optargex.Parse() {
   	  switch opt.ShortName {
   	  case "h":
   	  	// ...
   	  case "v":
   	  	// ...
   	  }
   }

 - Standard switch tokens are - and --. Can be modified by changing the vars
   optargex.ShortSwitch and optargex.LongSwitch.
 - Any arguments not associated with an option will be available in the
   optargex.Remainder slice after optargex.Parse() has been run.
 - Boolean flags require no value. The precense or absence of the flag is the
   value by itself. eg: flag '-n' is false if it's not found in os.Args, true if
   it is.
 - Exposes a Usage() function which prints options with their description and
   default values to the standard output. As opposed to the flag package, this
   outputs *neatly formatted* text. It prints output like the following. Note
   that this is the Usage() output of the options listed in optargex_test.go. It
   uses my own sexy multilineWrap() routine (see string.go):
 
--------------------------------------------------------------------------------
 Usage: ./6.out [options]:

 --source, -s: Path to the source folder. Here is some added description
               information which is completely useless, but it makes sure we can
               pimp our sexy Usage() output when dealing with lenghty, multi
               -line description texts.
    --bin, -b: Path to the binary folder.
   --arch, -a: Target architecture. (defaults to: amd64)
 --noproc, -n: Skip pre/post processing. (defaults to: false)
  --purge, -p: Clean compiled packages after linking is complete. (defaults to:
               false)
--------------------------------------------------------------------------------

 - We can organize options into 'sections', by interspersing their definition
   with header titles.
   
      optargex.Header("General options")
      optargex.Add("h", "help", "Displays this help.", false)
      optargex.Add("v", "version", "Displays version information.", false)

      optargex.Header("Find synonyms")
      optargex.Add("l", "list", "List all synonyms for the supplied words.", false)
      optargex.Add("d", "database", "Name of the thesaurus database to use.", v.DictDB)
      optargex.Add("s", "server", "Address of the thesaurus database to use.", v.DictServer)

   Since all options are printed in Usage() in the same order they are defined,
   this allows us to split the definitions up into useful and clearly defined
   categories. Making it easy for the reader to see which options are related to
   eachother. The standard Header output format "\n[%s]", which can be modified by
   optargs.HeaderFmt. This output of the example above is as follows:

   [General options]
        --help, -h: Displays this help.
     --version, -v: Displays version information.

   [Find synonyms]
        --list, -l: List all synonyms for the supplied words.
    --database, -d: Name of the thesaurus database to use. (defaults to: wn)
      --server, -s: Address of the thesaurus database to use. (defaults to:
                    dict.org:2628)


 - Specify arguments to options using whitespace.
 
   This is *not* accepted:
   -f=bar
   -fBar
  
   This *is* accepted:
   -f bar
   -f	"bar"

 - We support chaining shortform options together like this: -mnopq
   This considers each letter as a separate option. The last one can even have
   an argument if it needs one: -mnopq "foo"

--------------modify by wheelcomplex --------------

     optargex.SetVersion("some version information")
     
     add version information
     
     optargex.Version()
     
     show version information
     
     optargex.VersionString()
     
     get version information in strings

================================================================================
 USAGE
================================================================================

 Refer to doc/httpinfoserver.go for a usage example.
