package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
    if len(name) == 0 {
         return "One for you, one for me."
    }
       return "One for " + name + ", one for me."
}