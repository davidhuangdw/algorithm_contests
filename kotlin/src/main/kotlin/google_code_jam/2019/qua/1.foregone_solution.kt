package google_code_jam.`2019`.qua
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val T = readLine()!!.toInt()
    (1..T).forEach {
        val (N, ) = split_input()
        val A = mutableListOf<Char>()
        val B = mutableListOf<Char>()
        for(ch in N){
            if(ch == '4'){
                A.add('3')
                B.add('1')
            }else{
                A.add(ch)
                if(B.size > 0)
                    B.add('0')
            }
        }
        val a = A.joinToString("")
        val b = B.joinToString("")
        println("Case #${it}: $a $b")
    }
}
