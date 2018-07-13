package main

import(
  "fmt"
  //"github.com/rylans/getlang"
  "../getlang"
)

func main(){
    sentence := []string{ "My name is Bharat.", //English
    "मेरा नाम भारत है।", //Hindi
    "મારું નામ ભારત છે", //Gujrati
    "என் பெயர் பாரத்.", //Tamil
    "Mi nombre es Bharat.", //Spanish
    "Mon nom est Bharat.", //French
    "我的名字是巴拉特。", //Chinese
    "ਮੇਰਾ ਨਾਮ ਭਰਤ ਹੈ.", //punjabi
    "ನನ್ನ ಹೆಸರು ಭಾರತ್.", //Kannada
  }

     languages := map[string]string{
      "en": "English",
      "hi": "Hindi",
      "gu": "Gujrati",
      "ta": "Tamil",
      "es": "Spanish",
      "fr": "French",
      "zh": "Chinese",
      "pa": "Punjabi",
      "kn": "Kannada",
   }

   for i := 0; i< len(sentence); i++{
     language := getlang.FromString(sentence[i])
     fmt.Println("String==>",sentence[i])
     fmt.Printf("Language is %s and confidence is %f \n",languages[language.LanguageCode()],language.Confidence())
   }

}
