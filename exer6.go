package main


func camelcase(s string) int32 {
	count := int32(1) // Initialize count as int32
  
	for _, c := range s {
	  if int(c) >= 65 && int(c) <= 90 {
		count += 1
	  }
	}
  
	return count
}
func caesarCipher(s string, k int32) string {
    

}

func main(){
	camelcase("saveChangesInTheEditor")

}



