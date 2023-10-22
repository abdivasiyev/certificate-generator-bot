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

<pre><b>A.</b> %option_a%</pre>
<pre><b>B.</b> %option_b%</pre>
<pre><b>C.</b> %option_c%</pre>
<pre><b>D.</b> %option_d%</pre>`

	quizEndText = `Tabriklaymiz, siz barcha savollarni muvaffaqiyatli tugatdingiz!
<b>Tog'ri javoblar soni:</b> <pre>%corrects%</pre>
<b>Noto'g'ri javoblar soni:</b> <pre>%incorrects%</pre>`
)
