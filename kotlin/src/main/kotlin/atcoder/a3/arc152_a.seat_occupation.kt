package atcoder.a3

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (N, L) = input_ints()
    val A = input_ints()
    var s = 0
    val succ = A.withIndex().all{(i,a)->
//        println(listOf(i, a, L - (i+s)))
        (a<2 || (L - (i+s)) >= 2).also{
            s+=a
        }
    }
    println(if(succ) "Yes" else "No")
}