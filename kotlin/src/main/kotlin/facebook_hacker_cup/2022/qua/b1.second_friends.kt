
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (R, C) = input_ints()
        val V = (1..R).map{split_input()[0].toMutableList()}
        fun check(): String{
            if(R == 1 || C == 1){
                if(V.all{it.all{it == '.'}})
                    return "Possible\n" + (1..R).map{".".repeat(C)}.joinToString("\n")
                else return "Impossible"
            }
            return "Possible\n" + (1..R).map{"^".repeat(C)}.joinToString("\n")
        }
        println("Case #${it}: ${check()}")
    }
}
