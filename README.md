# voca

**V**ersatile **O**ffline **C**ommunication **A**ssistant — real-time translation TUI powered by local Ollama models.

```bash
go run . --model llama3.2:3b
```

The binary manages the full Ollama lifecycle on its own: starts the server if offline, pulls the model if missing, and cleans up on exit.

## Benchmark (EN → 14 languages)

Source text: "The quick brown fox jumps over the lazy dog. This sentence contains every letter of the English alphabet."

### Gemma 4 E2B QAT — default

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

### Phi-4 Mini 3.8B

| Lang | Time | Output |
|------|------|--------|
| it | 2.4s | Il veloce renato marrone salta sopra il cane pigro. Questa frase contiene ogni lettera dell'alfabeto inglese. |
| fr | 1.0s | Le renard brun rapide saute par-dessus le chien paresseux. Cette phrase contient chaque lettre de l'alphabet français. |
| de | 1.0s | Der schnelle braune Fuchs springt über den faulen Hund. Diese Aussage enthält jeden Buchstaben des englischen Alphabets. |
| es | 1.1s | El rápido zorro marrón salta sobre el perro perezoso. Esta frase contiene cada letra del alfabeto inglés. |
| pt | 1.2s | O rápido zorro marrón salta sobre el perro perezoso. Esta frase contiene cada letra del alfabeto inglés. |
| nl | 1.0s | De snelle bruine vos springt over de slappe hond. Deze zin bevat alle letters van het Engelse alfabet. |
| pl | 1.3s | Zawodny rudy lis skacze przez laźki psa. Ta zdanie zawiera każde literko angielskiego alfabetu. |
| ru | 1.4s | Крекша коричневый лис перепрыгивает через ленивую собаку. |
| ja | 1.6s | 速い茶色のキツネが怠け者の犬を飛び越えます。この文は英語アルファベットのすべての文字を含んでいます。 |
| zh | 1.0s | 快速的棕色狐狸跳过懒狗。这句话包含了英语字母表中的每一个字母。 |
| ko | 1.3s | 빠른 갈색 여우가 게으른 개를 뛰어넘는다. 이 문장은 영어 알파벳의 모든 글자를 포함하고 있다. |
| ar | 1.3s | الزرافة البرية السريعة الزرقاء تسبح فوق الكلب الكسول. |
| tr | 1.2s | Çok hızlı koyu kurt atlar ve yavaş köpekten geçer. |
| hi | 1.4s | जल्दी भुने कुत्ते ने आलसी बिल्ली को पार किया। |

**Average: ~1.3s** — fast as llama3.2, but quality degrades on non-European languages (Arabic, Hindi get wrong nouns).

### Nemotron 3 Nano 4B

| Lang | Time | Output |
|------|------|--------|
| it | 13.3s | Il veloce castano marrone salta sopra il cane pigro. Questa frase contiene ogni lettera dell'alfabeto inglese. |
| fr | 9.1s | Le rapide renard brun saute par-dessus le chien paresseux. Cette phrase contient toutes les lettres de l'alphabet anglais. |
| de | 8.3s | Der schnelle braune Fuchs springt über den faulen Hund. Dieser Satz enthält alle Buchstaben des englischen Alphabets. |
| es | 9.8s | El rápido zorro marrón salta sobre el perro perezoso. Esta oración contiene todas las letras del alfabeto inglés. |
| pt | 10.3s | O rabo rápido salta sobre o cão preguiçoso. Esta frase contém todas as letras do alfabeto inglês. |
| nl | 11.4s | De snelle bruine vos springt over de luipe hond. Deze zin bevat elke letter van het Engelse alfabet. |
| pl | 14.2s | Szybki brązowy lis przeskakuje nad leniwym psem. To zdanie zawiera wszystkie litery alfabetu angielskiego. |
| ru | 12.2s | Быстрая коричневая лиса прыгает над ленивым собакой. |
| ja | 14.2s | 素早い茶色の狐が怠け者の犬を飛び越える。この文には英語アルファベットのすべての文字が含まれている。 |
| zh | 8.9s | 快速的棕色狐狸跳过懒惰的狗。这句话包含了英语字母的全部字符。 |
| ko | 10.1s | 빠른 갈색 여우가 지저귀한 개를 뛰어넘는다. 이 문장은 영어 알파벳의 모든 글자를 포함한다. |
| ar | 8.6s | الذئب البني السريع يقفز فوق الكلب الكسول. هذه الجملة تحتوي على كل حرف من حروف اللغة الإنجليزية. |
| tr | 16.9s | Kısa mavi kedi, yorgulayıcı koşuyor. Bu cümle, İngilizce alfabesinin her harfi içerir. |
| hi | 12.8s | तेज़ नरम बैंगल फोक ने लाजी कुकी को पार कर रहा है। |

**Average: ~11.4s** — quality close to Gemma 4.

### LFM 2.5 1.2B

| Lang | Time | Output |
|------|------|--------|
| it | 10.1s | Il veloce volpe marrone salta sopra il cane pigro. Questa frase contiene ogni lettera dell'alfabeto inglese. |
| fr | 0.5s | Le rapide renard brun saute par-dessus le chien paresseux. Cette phrase contient chaque lettre de l'alphabet anglais. |
| de | 0.6s | Die schnelle braune Fuchs springt über den faulen Hund. Dieser Satz enthält jedes Buchstaben des englischen Alphabets. |
| es | 0.5s | El rápido zorro marrón salta sobre el perro perezoso. Esta frase contiene cada letra del inglés. |
| pt | 0.7s | A raposa rápida salta sobre o cão preguiçoso. Esta frase contém todas as letras do alfabeto inglês. |
| nl | 0.7s | De snelle bruine zeehond springt over de laaih hond. |
| pl | 0.7s | Szybki pies zimny wskakuje nad lazującego psa. |
| ru | 0.7s | Скороковый медведий оскользнуется над лениным собаку. |
| ja | 0.3s | この文には英語のすべての文字が含まれています。 |
| zh | 0.4s | 这句话包含了英文字母的每一个字符。 |
| ko | 0.7s | 이 빠른 갈색 여우가 게으른 개 위를 뛰어넘습니다. 이 문장은 영문 알파벳의 모든 글자를 포함하고 있습니다. |
| ar | 0.8s | الثعلب الأسود السري يجتاز الكلب النائم. الجملة تحتوي على كل حرف من حروف الأبجدية الإنجليزية. |
| tr | 1.0s | Kısa bir kızıl gürüş koyduğu gözlemleniz. |
| hi | 1.8s | ते सुप्य भारी खरगोल को ऊपर चम्मच जाता है। |

**Average: ~1.4s** (cold start IT: 10s, excluding it: ~0.7s) — fastest and lightest (1.2 GB). Limited to 8 official languages, produces nonsense on unsupported ones.

## Comparison

| Model | Params | Size | Avg time | Best for |
|-------|--------|------|----------|----------|
| Gemma 4 E2B QAT | 4.6B | 4.3 GB | ~11.2s | Best quality on all 25 languages |
| Nemotron 3 Nano 4B | 30B MoE (3.5B active) | 2.8 GB | ~11.4s | Good quality, large context (256K) |
| Phi-4 Mini | 3.8B | 2.5 GB | ~1.3s | Fast, good for European langs |
| Llama 3.2 3B | 3B | 2.0 GB | ~1.1s | Fastest, decent for EU langs |
| LFM 2.5 1.2B | 1.2B | 1.2 GB | ~0.7s | Ultra-light, only 8 langs |

## Model selection

```bash
go run .

# Fast, good for European languages
go run . --model llama3.2:3b

# Use a specific model
go run . --model <model-name>
```
