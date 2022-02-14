package constants

import "strconv"

/*
	Bu PL/SQL sorgusu önce kullanıcı varmı? yokmu? diye arar(bunu rec değişkeninin değerine bakarak yapar)
	daha sonra var ise var olan kullanıcının cüzdan değerini günceller
	yok ise initial değer cüzdan oluşturur.
*/

func UpdateBalanceQuery(username string, values ...int) string {
	return "DECLARE rec NUMBER;" + "\n" +
		"BEGIN" + "\n" +
		"SELECT COUNT(*) INTO rec FROM Wallet WHERE username ='" + username + "'  GROUP BY username;" + "\n" +
		"IF rec > 0 THEN" + "\n" +
		"UPDATE Wallet" + "\n" +
		"SET balanceAmount = " + strconv.Itoa(values[0]) + "\n" +
		"WHERE username = '" + username + "';" + "\n" +
		"ELSE " + "\n" +
		"INSERT INTO" + "\n" +
		"Wallet (username, balanceamount)" + "\n" +
		"VALUES" + "\n" +
		"('" + username + "', " + strconv.Itoa(values[1]) + ");" + "\n" +
		"END IF;" + "\n" +
		"EXCEPTION" + "\n" +
		"WHEN NO_DATA_FOUND THEN" + "\n" +
		"rec := NULL;" + "\n" +
		"INSERT INTO" + "\n" +
		"Wallet (username, balanceamount)" + "\n" +
		"VALUES" + "\n" +
		"('" + username + "', " + strconv.Itoa(values[1]) + ");" + "\n" +
		"END;"
}

/*
	Bu PL/SQL sorgusu girilen değere göre cüzdan değerini artırır ya da azaltır.
*/

func PostBalanceQuery(username string, value int) string {
	return "DECLARE rec NUMBER;" + "\n" +
		"BEGIN" + "\n" +
		"SELECT COUNT(*) INTO rec FROM Wallet WHERE username ='" + username + "'  GROUP BY username;" + "\n" +
		"IF rec > 0 THEN" + "\n" +
		"UPDATE Wallet" + "\n" +
		"SET balanceAmount = balanceAmount+" + strconv.Itoa(value) + "\n" +
		"WHERE username = '" + username + "';" + "\n" +
		"END IF;" + "\n" +
		"EXCEPTION" + "\n" +
		"WHEN NO_DATA_FOUND THEN" + "\n" +
		"rec := NULL;" + "\n" +
		"END;"
}
