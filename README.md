# voca

**V**ersatile **O**ffline **C**ommunication **A**ssistant — real-time translation TUI powered by local Ollama models.

```bash
go run . --model llama3.2:3b
```

The binary manages the full Ollama lifecycle on its own: starts the server if offline, pulls the model if missing, and cleans up on exit.

## Benchmark (EN → 14 languages)

Source text: "The quick brown fox jumps over the lazy dog. This sentence contains every letter of the English alphabet."

### Gemma 4 E2B QAT (4.6B) — default

| Lang | Time | Output |
|------|------|--------|
| it | 7.7s | La volpe marrone veloce salta sopra il cane pigro. Questa frase contiene ogni lettera dell'alfabeto inglese. |
| fr | 7.5s | Le renard brun rapide saute par-dessus le chien paresseux. Cette phrase contient chaque lettre de l'alphabet anglais. |
| de | 7.9s | Der schnelle braune Fuchs springt über den faulen Hund. Dieser Satz enthält jeden Buchstaben des englischen Alphabets. |
| es | 7.1s | El zorro marrón rápido salta sobre el perro perezoso. Esta oración contiene cada letra del alfabeto inglés. |
| pt | 7.5s | A raposa marrom rápida pula sobre o cão preguiçoso. Esta frase contém cada letra do alfabeto inglês. |
| nl | 8.0s | De snelle bruine vos springt over de luie hond. Deze zin bevat elke letter van het Engelse alfabet. |
| pl | 5.9s | Szybki brązowy lis przeskakuje nad leniwym psem. To zdanie zawiera każdą literę angielskiego alfabetu. |
| ru | 9.0s | Быстрая коричневая лиса перепрыгивает через ленивую собаку. Это предложение содержит каждую букву английского алфавита. |
| ja | 15.0s | 素早い茶色の狐が怠け者の犬を飛び越えます。この文は英語のアルファベットのすべての文字を含んでいます。 |
| zh | 9.0s | 敏捷的棕色狐狸跳过了懒狗。这句话包含了英文字母表中的每一个字母。 |
| ko | 14.3s | 빠른 갈색 여우가 게으른 개를 뛰어넘습니다. 이 문장은 영어 알파벳의 모든 글자를 포함합니다. |
| ar | 15.2s | يقفز الثعلب البني السريع فوق الكلب الكسول. تحتوي هذه الجملة على كل حرف من الحروف الأبجدية الإنجليزية. |
| tr | 10.1s | Hızlı kahverengi tilki tembel köpeğin üzerinden atlar. Bu cümle İngiliz alfabesindeki her harfi içerir. |
| hi | 27.3s | तेज़ भूरी लोमड़ी आलसी कुत्ते के ऊपर से कूदती है। इस वाक्य में अंग्रेज़ी वर्णमाला का हर अक्षर है। |

**Average: ~11.2s**

### Llama 3.2 3B

| Lang | Time | Output |
|------|------|--------|
| it | 3.1s | Il volpe rosso veloce salta sopra il cane sonnolento. Questa frase contiene ogni lettera dell'alfabeto inglese. |
| fr | 1.1s | Le renard brun rapide saute au-dessus du chien paresseux. Cette phrase contient tous les lettres de l'alphabet français. |
| de | 1.1s | Der schnelle braune Fuchs springt über den müde hörigen Hund. Diese Sätze enthält jede Buchstabe des deutschen Alphabets. |
| es | 1.1s | El zorro marrón rápido salta sobre el perro perezoso. Esta oración contiene todos los letras del alfabeto inglés. |
| pt | 1.1s | O corrente vistoso vulpino salta sobre o molez do gato. Esta frase contém todos os letras do alfabeto português. |
| nl | 1.0s | De snelle bruine vos springt over de slaapdog. Deze zin bevat elke letter van het Engelse alfabet. |
| pl | 0.7s | Kutty bruny wół skacze nad słabego psa. |
| ru | 1.0s | Скорый коричневый Kits jump over lazy dog. Этот предложение содержит каждую букву английского алфавита. |
| ja | 1.2s | はいくみつう すばらしい こぶし つりが うるうの ひどい ねこを こわして いる。 |
| zh | 0.9s | 快乐的棕色狐狸跳过懒狗。这个句子包含了所有英文字母。 |
| ko | 1.2s | FOX가 빠르게 갈색의 개를 넘기고 수동한 개가 있습니다. 이 문장에는 영어 알파벳이 모든 lettre가 포함됩니다. |
| ar | 0.7s | الزاحف الأصفر الحلو يطول على السنّاه. |
| tr | 1.2s | Kısa kahverengi fox, uyuşuk köpeği över atar. Bu cümle, İngilizce alfabeinin her bir harfini içerir. |
| hi | 1.7s | तेज़ काला वolf जंप करता है सुन्न कुत्ते को पार करता है। यह वाक्य अंग्रेजी भाषा में हर एक अक्षर को शामिल करता है। |

**Average: ~1.1s** — very fast, quality degrades on non-European languages.

## Model selection

```bash
go run .

# Fast, good for European languages
go run . --model llama3.2:3b
```
