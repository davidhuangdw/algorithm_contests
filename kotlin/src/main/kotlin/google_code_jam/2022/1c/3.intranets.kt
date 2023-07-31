package google_code_jam.`2022`.`1c`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        // todo it: https://codingcompetitions.withgoogle.com/codejam/round/0000000000877b42/0000000000afeb38
        println("Case #${it}: ")
    }
}
