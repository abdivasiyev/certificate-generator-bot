package bot

const (
	userCreatedTextOnStart = `Assalomu alaykum! 👋
Ushbu bot orqali siz Go dasturlash tiliga oid test savollaridan o'tishingiz mumkin.
Testni boshlash uchun quyidagi tugmani bosing!
👇👇👇👇👇👇👇👇👇`

	userGotTextOnStart = `Assalomu alaykum! 👋
Yana ko'rishib turganimdan xursandman!
Ushbu bot orqali siz Go dasturlash tiliga oid test savollaridan o'tishingiz mumkin.
<b>Hurmatli foydanaluvchi, siz hozir %quiz_id% raqamli savoldasiz.</b>
Testni davom ettirish uchun quyidagi tugmani bosing!
👇👇👇👇👇👇👇👇👇`

	quizText = `%quiz_id%. %quiz_question%
<pre>A. %option_a%</pre>
<pre>B. %option_b%</pre>
<pre>C. %option_c%</pre>
<pre>D. %option_d%</pre>`

	quizEndText = `Tabriklaymiz, siz barcha savollarni muvaffaqiyatli tugatdingiz!
Tog'ri javoblar soni: <pre>%corrects%</pre>
Noto'g'ri javoblar soni: <pre>%incorrects%</pre>`
)
